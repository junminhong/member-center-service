package usecase

import (
	"github.com/junminhong/member-center-service/domain"
	"github.com/junminhong/member-center-service/pkg/jwt"
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
	"go.uber.org/zap"
	"strings"
	"time"
)

type authUseCase struct {
	AuthRepo   domain.AuthRepository
	MemberRepo domain.MemberRepository
	ZapLogger  *zap.SugaredLogger
}

func NewAuthUseCase(authRepo domain.AuthRepository, memberRepo domain.MemberRepository, logger *zap.SugaredLogger) domain.AuthUseCase {
	return &authUseCase{authRepo, memberRepo, logger}
}

func (authUseCase *authUseCase) AuthEmailToken(request requester.AuthEmailToken) responser.Response {
	email := authUseCase.AuthRepo.GetEmailByEmailToken(request.EmailToken)
	if email == "" {
		// Email token 過期
		return responser.Response{
			ResultCode: responser.EmailTokenExpiredErr.Code(),
			Message:    responser.EmailTokenExpiredErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	member := authUseCase.MemberRepo.GetMemberByEmail(email)
	if member.Email == "" {
		return responser.Response{
			ResultCode: responser.MemberNotFoundErr.Code(),
			Message:    responser.MemberNotFoundErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	member.EmailAuth = true
	authUseCase.MemberRepo.SaveMember(member)
	return responser.Response{
		ResultCode: responser.EmailAuthOk.Code(),
		Message:    responser.EmailAuthOk.Message(),
		Data:       "",
		TimeStamp:  time.Now(),
	}
}

func (authUseCase *authUseCase) RefreshAtomicToken(refreshAtomicToken string) responser.Response {
	if !jwt.VerifyAtomicToken(refreshAtomicToken) {
		//過期了
		return responser.Response{
			ResultCode: responser.AtomicTokenExpiredErr.Code(),
			Message:    responser.AtomicTokenExpiredErr.Reload("Refresh Atomic Token已經過期了").Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	memberUUID := authUseCase.MemberRepo.GetMemberUUIDByAtomicToken(refreshAtomicToken)
	if memberUUID == "" {
		return responser.Response{
			ResultCode: responser.AtomicTokenExpiredErr.Code(),
			Message:    responser.AtomicTokenExpiredErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	member := authUseCase.MemberRepo.GetMemberByUUID(memberUUID)
	if strings.Compare(member.RefreshAtomicToken, refreshAtomicToken) != 0 {
		return responser.Response{
			ResultCode: responser.NotAtomicTokenErr.Code(),
			Message:    responser.NotAtomicTokenErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		}
	}
	authUseCase.AuthRepo.ClearAtomicToken(member.AtomicToken)
	atomicToken := jwt.GenerateAtomicToken(4)
	member.AtomicToken = atomicToken
	authUseCase.MemberRepo.SaveMember(member)
	authUseCase.MemberRepo.StoreAtomicToken(atomicToken, member.UUID, 4)
	return responser.Response{
		ResultCode: responser.RefreshAtomicTokenOk.Code(),
		Message:    responser.RefreshAtomicTokenOk.Message(),
		Data: responser.RefreshAtomicToken{
			AtomicToken: atomicToken,
		},
		TimeStamp: time.Now(),
	}
}
