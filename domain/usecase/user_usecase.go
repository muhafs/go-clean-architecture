package usecase

import (
	"context"

	"github.com/muhafs/go-clean-architecture/domain/entity"
)

type SignupUsecase interface {
	Create(c context.Context, user *entity.User) error
	FindByEmail(email string) (entity.User, error)
	CreateAccessToken(user *entity.User, secret string, expiry int) (string, error)
	CreateRefreshToken(user *entity.User, secret string, expiry int) (string, error)
}
