package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/api/route"
	"github.com/muhafs/go-clean-architecture/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mysql
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(env, db, timeout, gin)

	gin.Run(":3000")
}
