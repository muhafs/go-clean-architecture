package route

import (
	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/api/controller"
	"github.com/muhafs/go-clean-architecture/bootstrap"
	"github.com/muhafs/go-clean-architecture/repository"
	"github.com/muhafs/go-clean-architecture/usecase"
	"gorm.io/gorm"
)

func NewCategoryRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUscs := usecase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := controller.CategoryController{
		CategoryUsecase: categoryUscs,
		Env:             env,
	}

	group.POST("/category", categoryCtrl.Create)
}
