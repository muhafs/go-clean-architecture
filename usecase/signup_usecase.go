package usecase

import (
	"context"
	"time"

	"github.com/muhafs/go-clean-architecture/domain/entity"
	"github.com/muhafs/go-clean-architecture/domain/repository"
	"github.com/muhafs/go-clean-architecture/domain/usecase"
)

type SignupUsecase struct {
	UserRepository repository.UserRepository
	ContextTimeout time.Duration
}

func NewSignupUsecase(repo repository.UserRepository, timeout time.Duration) usecase.SignupUsecase {
	return &SignupUsecase{
		UserRepository: repo,
		ContextTimeout: timeout,
	}
}

func (su *SignupUsecase) Create(c context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(c, su.ContextTimeout)
	defer cancel()

	return su.UserRepository.Create(ctx, user)
}

func (su *SignupUsecase) FindByEmail(email string) (entity.User, error) {
	return su.UserRepository.FindByEmail(email)
}

func (su *SignupUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (string, error) {
	return "", nil
}

func (su *SignupUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (string, error) {
	return "", nil
}
