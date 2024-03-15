package repository

import (
	"context"

	"github.com/muhafs/go-clean-architecture/domain/entity"
)

type UserRepository interface {
	Create(c context.Context, user *entity.User) error
	FindAll() ([]entity.User, error)
	FindOne(id int) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
