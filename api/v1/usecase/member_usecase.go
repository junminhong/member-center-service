package usecase

import (
	"github.com/google/uuid"
	"github.com/junminhong/member-center-service/domain"
	"github.com/junminhong/member-center-service/pkg/encode"
	"github.com/junminhong/member-center-service/pkg/gcp"
	"github.com/junminhong/member-center-service/pkg/jwt"
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
	"github.com/junminhong/member-center-service/pkg/smtp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mime/multipart"
	"strings"
	"time"
)

type memberUseCase struct {
	MemberRepo domain.MemberRepository
	ZapLogger  *zap.SugaredLogger
}

func NewMemberUseCase(memberRepository domain.MemberRepository, logger *zap.SugaredLogger) domain.MemberUseCase {
	return &memberUseCase{memberRepository, logger}
}

func (memberUseCase *memberUseCase) Register(request requester.Register) responser.Response {
	member := memberUseCase.MemberRepo.GetMemberByEmail(request.Email)
	if request.Email == member.Email {
		return responser.Response{
			ResultCode: responser.EmailExistsErr.Code(),
			Message:    responser.EmailExistsErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	member.UUID = uuid.NewString()
	member.Email = request.Email
	member.Password = request.Password
	member.MemberInfo = domain.MemberInfo{NickName: request.NickName}
	if !memberUseCase.MemberRepo.Register(member) {
		return responser.Response{
			ResultCode: responser.MemberRegisterErr.Code(),
			Message:    responser.MemberRegisterErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	emailToken := encode.EmailToken(member.Email)
	var smtp = &smtp.EmailMsg{}
	smtp.SetUpEmailMsg("帳戶認證信件", "請點擊以下網址連結："+viper.GetString("APP.AUTH_EMAIL_PATH")+"?email_token="+emailToken)
	go smtp.SendEmail(member.Email)
	go memberUseCase.MemberRepo.StoreEmailToken(emailToken, member.Email, 60)
	return responser.Response{
		ResultCode: responser.MemberRegisterOk.Code(),
		Message:    responser.MemberRegisterOk.Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (memberUseCase *memberUseCase) Login(request requester.Login) responser.Response {
	member := memberUseCase.MemberRepo.GetMemberByEmail(request.Email)
	if member.Email == "" {
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	if strings.Compare(member.Password, request.Password) != 0 {
		return responser.Response{
			ResultCode: responser.PasswordNotMatchErr.Code(),
			Message:    responser.PasswordNotMatchErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	if !member.EmailAuth {
		return responser.Response{
			ResultCode: responser.EmailNotAuthErr.Code(),
			Message:    responser.EmailNotAuthErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	atomicToken := jwt.GenerateAtomicToken(4)
	refreshAtomicToken := jwt.GenerateAtomicToken(24)
	member.AtomicToken = atomicToken
	member.RefreshAtomicToken = refreshAtomicToken
	memberUseCase.MemberRepo.StoreAtomicToken(atomicToken, member.UUID, 4)
	memberUseCase.MemberRepo.StoreAtomicToken(refreshAtomicToken, member.UUID, 24)
	memberUseCase.MemberRepo.SaveMember(member)
	return responser.Response{
		ResultCode: responser.LoginOk.Code(),
		Message:    responser.LoginOk.Message(),
		Data: responser.Login{
			AtomicToken:        atomicToken,
			RefreshAtomicToken: refreshAtomicToken,
		},
		TimeStamp: time.Now(),
	}
}

func (memberUseCase *memberUseCase) ResendEmail(request requester.ResendEmail) responser.Response {
	member := memberUseCase.MemberRepo.GetMemberByEmail(request.Email)
	if member.Email == "" {
		// 查無此人
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	emailToken := encode.EmailToken(request.Email)
	var smtp = &smtp.EmailMsg{}
	smtp.SetUpEmailMsg("帳戶認證信件", "請點擊以下網址連結："+viper.GetString("APP.AUTH_EMAIL_PATH")+"?email_token="+emailToken)
	go smtp.SendEmail(member.Email)
	go memberUseCase.MemberRepo.StoreEmailToken(emailToken, member.Email, 60)
	return responser.Response{
		ResultCode: responser.ResendEmailOk.Code(),
		Message:    responser.ResendEmailOk.Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (memberUseCase *memberUseCase) ResetPassword(request requester.ResetPassword, atomicToken string) responser.Response {
	memberUUID, response := memberUseCase.checkAtomicToken(atomicToken)
	if memberUUID == "" {
		return response
	}
	member := memberUseCase.MemberRepo.GetMemberByUUID(memberUUID)
	if strings.Compare(request.OldPassword, member.Password) != 0 {
		// 舊密碼輸入錯誤
		return responser.Response{
			ResultCode: responser.PasswordNotMatchErr.Code(),
			Message:    responser.PasswordNotMatchErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	member.Password = request.NewPassword
	memberUseCase.MemberRepo.SaveMember(member)
	return responser.Response{
		ResultCode: responser.ResetPasswordOk.Code(),
		Message:    responser.ResetPasswordOk.Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (memberUseCase *memberUseCase) ForgetPassword(request requester.ForgetPassword) responser.Response {
	member := memberUseCase.MemberRepo.GetMemberByEmail(request.Email)
	if member.Email == "" {
		// 查無此人
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	var smtp = &smtp.EmailMsg{}
	smtp.SetUpEmailMsg("忘記用戶密碼", "您的密碼是："+member.Password+"，貼心提醒！建議重置密碼，可以有效保障帳戶安全性哦。")
	go smtp.SendEmail(member.Email)
	return responser.Response{
		ResultCode: responser.ResendEmailOk.Code(),
		Message:    responser.ResendEmailOk.Reload("已寄送忘記密碼信").Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (memberUseCase *memberUseCase) GetProfile(atomicToken string) responser.Response {
	memberUUID, response := memberUseCase.checkAtomicToken(atomicToken)
	if memberUUID == "" {
		return response
	}
	member := memberUseCase.MemberRepo.GetMemberByUUID(memberUUID)
	if member.Email == "" {
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	memberInfo := memberUseCase.MemberRepo.GetMemberInfo(member)
	mugShotPath := gcp.GetFileForGCS(memberInfo.MugShotPath)
	return responser.Response{
		ResultCode: responser.GetProfileOk.Code(),
		Message:    responser.GetProfileOk.Message(),
		Data: responser.GetProfile{
			NickName:    memberInfo.NickName,
			MugShotPath: mugShotPath,
		},
		TimeStamp: time.Now(),
	}
}

func (memberUseCase *memberUseCase) EditProfile(atomicToken string, request requester.EditProfile) responser.Response {
	memberUUID, response := memberUseCase.checkAtomicToken(atomicToken)
	if memberUUID == "" {
		return response
	}
	member := memberUseCase.MemberRepo.GetMemberByUUID(memberUUID)
	if member.Email == "" {
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	memberInfo := memberUseCase.MemberRepo.GetMemberInfo(member)
	memberInfo.NickName = request.NickName
	memberUseCase.MemberRepo.SaveMemberInfo(memberInfo)
	return responser.Response{
		ResultCode: responser.EditProfileOk.Code(),
		Message:    responser.EditProfileOk.Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (memberUseCase *memberUseCase) UploadMugShot(atomicToken string, file multipart.File, uploadFile *multipart.FileHeader) responser.Response {
	memberUUID, response := memberUseCase.checkAtomicToken(atomicToken)
	if memberUUID == "" {
		return response
	}
	member := memberUseCase.MemberRepo.GetMemberByUUID(memberUUID)
	if member.Email == "" {
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	memberInfo := memberUseCase.MemberRepo.GetMemberInfo(member)
	if (uploadFile.Size / 1000000) > 10 {
		return responser.Response{
			ResultCode: responser.UploadFileTooBigErr.Code(),
			Message:    responser.UploadFileTooBigErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	fileType := strings.Split(uploadFile.Header.Get("Content-Type"), "/")
	if fileType[0] != "image" {
		return responser.Response{
			ResultCode: responser.UploadFileNotImgErr.Code(),
			Message:    responser.UploadFileNotImgErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	if !gcp.InsertFileToGCS("mug-shot/", member.UUID, uploadFile, file) {
		return responser.Response{
			ResultCode: responser.UploadFileErr.Code(),
			Message:    responser.UploadFileErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	memberInfo.MugShotPath = "mug-shot/" + member.UUID
	memberUseCase.MemberRepo.SaveMemberInfo(memberInfo)
	return responser.Response{
		ResultCode: responser.UploadFileOk.Code(),
		Message:    responser.UploadFileOk.Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (memberUseCase *memberUseCase) checkAtomicToken(atomicToken string) (memberUUID string, response responser.Response) {
	if !jwt.VerifyAtomicToken(atomicToken) {
		// atomic token過期了
		return "", responser.Response{
			ResultCode: responser.AtomicTokenExpiredErr.Code(),
			Message:    responser.AtomicTokenExpiredErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	memberUUID = memberUseCase.MemberRepo.GetMemberUUIDByAtomicToken(atomicToken)
	if memberUUID == "" {
		return "", responser.Response{
			ResultCode: responser.AtomicTokenExpiredErr.Code(),
			Message:    responser.AtomicTokenExpiredErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	return memberUUID, responser.Response{}
}
