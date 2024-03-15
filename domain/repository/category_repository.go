package repository

import "github.com/muhafs/go-clean-architecture/domain/entity"

type CategoryRepository interface {
	Create(ctg *entity.Category) error
	FindAll() ([]entity.Category, error)
	FindOne(id int) (entity.Category, error)
}
