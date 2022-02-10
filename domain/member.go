package domain

import (
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type Member struct {
	gorm.Model
	ID                 int    `gorm:"primaryKey"`
	UUID               string `gorm:"index;unique"`
	Email              string `gorm:"unique"`
	Password           string
	AtomicToken        string
	RefreshAtomicToken string
	EmailAuth          bool
	ThirdAuthToken     string
	SafePassword       string
	AtomicPoint        int
	MemberInfo         MemberInfo `gorm:"foreignKey:MemberInfoID"`
	ActivatedAt        time.Time
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

type MemberInfo struct {
	gorm.Model
	MemberInfoID int `gorm:"primaryKey"`
	NickName     string
	MugShotPath  string
	SocialInfo   SocialInfo `gorm:"foreignKey:SocialInfoID"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}

type SocialInfo struct {
	gorm.Model
	SocialInfoID int `gorm:"primaryKey"`
	SocialType   string
	SocialUrl    string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type MemberRepository interface {
	GetMemberByEmail(email string) (member Member)
	GetMemberByAtomicToken(atomicToken string) (member Member)
	GetMemberInfo(member Member) (memberInfo MemberInfo)
	GetMemberUUIDByAtomicToken(atomicToken string) (memberUUID string)
	GetMemberByUUID(memberUUID string) (member Member)
	Register(member Member) bool
	StoreEmailToken(token string, email string, timeLimit int)
	StoreAtomicToken(token string, memberUUID string, timeLimit int)
	SaveMember(member Member)
	SaveMemberInfo(memberInfo MemberInfo)
}

type MemberUseCase interface {
	Register(request requester.Register) responser.Response
	Login(request requester.Login) responser.Response
	ResendEmail(request requester.ResendEmail) responser.Response
	ResetPassword(request requester.ResetPassword, atomicToken string) responser.Response
	ForgetPassword(request requester.ForgetPassword) responser.Response
	GetProfile(atomicToken string) responser.Response
	EditProfile(atomicToken string, request requester.EditProfile) responser.Response
	UploadMugShot(atomicToken string, file multipart.File, uploadFile *multipart.FileHeader) responser.Response
}
