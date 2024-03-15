package usecase

import (
	"github.com/muhafs/go-clean-architecture/domain/entity"
	"github.com/muhafs/go-clean-architecture/domain/repository"
	"github.com/muhafs/go-clean-architecture/domain/usecase"
)

type CategoryUsecase struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) usecase.CategoryUsecase {
	return &CategoryUsecase{
		CategoryRepository: repo,
	}
}

func (cu *CategoryUsecase) Create(ctg *entity.Category) error {
	return cu.CategoryRepository.Create(ctg)
}

func (cu *CategoryUsecase) FindAll() ([]entity.Category, error) {
	return cu.CategoryRepository.FindAll()
}

func (cu *CategoryUsecase) FindOne(id int) (entity.Category, error) {
	return cu.CategoryRepository.FindOne(id)
}
