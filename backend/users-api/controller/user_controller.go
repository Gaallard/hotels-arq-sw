package users

import (
	"net/http"
	"strconv"
	dtoUsers "users-api/dto"
	users "users-api/service"

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

func GetUserById(context *gin.Context) {
	idParam := context.Param("idUser")
	idUser, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	response, apiErr := users.GetHotelById(idUser)
	if apiErr != nil {
		context.JSON(apiErr.Status(), apiErr)
		return
	}
	context.JSON(http.StatusOK, response)
}
