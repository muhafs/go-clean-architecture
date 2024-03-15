package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/api/middleware"
	"github.com/muhafs/go-clean-architecture/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	// public APIs
	publicRouter := gin.Group("")
	NewSignupRouter(env, db, timeout, publicRouter)

	// private APIs
	privateRouter := gin.Group("")
	privateRouter.Use(middleware.AuthMiddleware(env.AccessTokenSecret))
	NewCategoryRouter(env, db, privateRouter)
}
