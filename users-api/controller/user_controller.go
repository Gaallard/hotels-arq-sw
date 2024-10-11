package users

import (
	"net/http"
	dtoUsers "users-api/users-api/dto"
	"users-api/users-api/service"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var userRequest dtoUsers.UserDto
	context.BindJSON(&userRequest)
	response, err := users.RegisterUser(userRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, response)
}

func Login(context *gin.Context) {
	var loginRequest dtoUsers.UserDto
	context.BindJSON(&loginRequest)
	response, err := users.Login(loginRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}
