package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/junminhong/member-center-service/domain"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type memberRepo struct {
	DB        *gorm.DB
	Redis     *redis.Client
	ZapLogger *zap.SugaredLogger
}

func NewMemberRepo(db *gorm.DB, redis *redis.Client, logger *zap.SugaredLogger) domain.MemberRepository {
	return &memberRepo{db, redis, logger}
}

func (memberRepo *memberRepo) Register(member domain.Member) bool {
	err := memberRepo.DB.Create(&member).Error
	if err != nil {
		return false
	}
	return true
}

func (memberRepo *memberRepo) GetMemberByEmail(email string) (member domain.Member) {
	memberRepo.DB.Where("email = ?", email).First(&member)
	return member
}

func (memberRepo *memberRepo) GetMemberByAtomicToken(atomicToken string) (member domain.Member) {
	value := memberRepo.Redis.Get(context.Background(), atomicToken).Val()
	if value != "" {
		memberRepo.DB.Where("uuid = ?", value).First(&member)
	}
	return member
}

func (memberRepo *memberRepo) GetMemberByUUID(memberUUID string) (member domain.Member) {
	memberRepo.DB.Where("uuid = ?", memberUUID).First(&member)
	return member
}

func (memberRepo *memberRepo) GetMemberInfo(member domain.Member) (memberInfo domain.MemberInfo) {
	memberRepo.DB.Model(&member).Association("MemberInfo").Find(&memberInfo)
	return memberInfo
}

func (memberRepo *memberRepo) GetMemberUUIDByAtomicToken(atomicToken string) (memberUUID string) {
	memberUUID = memberRepo.Redis.Get(context.Background(), atomicToken).Val()
	return memberUUID
}

func (memberRepo *memberRepo) StoreEmailToken(token string, email string, timeLimit int) {
	memberRepo.Redis.Set(context.Background(), token, email, time.Duration(timeLimit)*time.Minute).Err()
}

func (memberRepo *memberRepo) StoreAtomicToken(token string, memberUUID string, timeLimit int) {
	memberRepo.Redis.Set(context.Background(), token, memberUUID, time.Duration(timeLimit)*time.Hour).Err()
}

func (memberRepo *memberRepo) SaveMember(member domain.Member) {
	memberRepo.DB.Save(member)
}

func (memberRepo *memberRepo) SaveMemberInfo(memberInfo domain.MemberInfo) {
	memberRepo.DB.Save(memberInfo)
}
