package repository

import (
	"context"

	"github.com/muhafs/go-clean-architecture/domain/entity"
	"github.com/muhafs/go-clean-architecture/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) Create(c context.Context, user *entity.User) error {
	err := ur.DB.WithContext(c).Create(user).Error
	return err
}

func (ur *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := ur.DB.Find(&users).Error

	return users, err
}

func (ur *UserRepository) FindOne(id int) (entity.User, error) {
	var user entity.User
	err := ur.DB.First(&user, id).Error

	return user, err
}

func (ur *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := ur.DB.First(&user, "email = ?", email).Error

	return user, err
}
