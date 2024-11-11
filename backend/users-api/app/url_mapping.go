package app

import (
	usersController "backend/controllers/users"

	"backend/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func mapUrls(router *gin.Engine) {

	router.POST("/users", usersController.UsuarioInsert)
	router.POST("/users/login", usersController.Login)
	router.GET("/users/token", usersController.Extrac)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/users", usersController.GetUserByName)

	}
	log.Info("Finishing mappings configurations")

}
