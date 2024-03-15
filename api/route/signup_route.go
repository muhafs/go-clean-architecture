package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/api/controller"
	"github.com/muhafs/go-clean-architecture/bootstrap"
	"github.com/muhafs/go-clean-architecture/repository"
	"github.com/muhafs/go-clean-architecture/usecase"
	"gorm.io/gorm"
)

func NewSignupRouter(env *bootstrap.Env, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)
	signupUscs := usecase.NewSignupUsecase(userRepo, timeout)
	signupCtrl := controller.SignupController{
		SignupUsecase: signupUscs,
		Env:           env,
	}

	group.POST("/signup", signupCtrl.Signup)
}
