package domain

import (
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
)

type AuthRepository interface {
	GetEmailByEmailToken(emailToken string) (email string)
	ClearAtomicToken(atomicToken string)
}

type AuthUseCase interface {
	AuthEmailToken(request requester.AuthEmailToken) responser.Response
	RefreshAtomicToken(refreshAtomicToken string) responser.Response
}
