package repository

import (
	"github.com/muhafs/go-clean-architecture/domain/entity"
	"github.com/muhafs/go-clean-architecture/domain/repository"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (cr *CategoryRepository) Create(ctg *entity.Category) error {
	err := cr.DB.Create(ctg).Error
	return err
}

func (cr *CategoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := cr.DB.Find(&categories).Error

	return categories, err
}

func (cr *CategoryRepository) FindOne(id int) (entity.Category, error) {
	var category entity.Category
	err := cr.DB.First(&category, id).Error

	return category, err
}
