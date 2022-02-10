package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/junminhong/member-center-service/domain"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type authRepo struct {
	DB        *gorm.DB
	Redis     *redis.Client
	ZapLogger *zap.SugaredLogger
}

func NewAuthRepo(db *gorm.DB, redis *redis.Client, logger *zap.SugaredLogger) domain.AuthRepository {
	return &authRepo{db, redis, logger}
}

func (authRepo *authRepo) GetEmailByEmailToken(emailToken string) (email string) {
	email = authRepo.Redis.Get(context.Background(), emailToken).Val()
	return email
}

func (authRepo *authRepo) ClearAtomicToken(atomicToken string) {
	_, err := authRepo.Redis.Get(context.Background(), atomicToken).Result()
	if err != redis.Nil {
		authRepo.Redis.Del(context.Background(), atomicToken)
	}
}
