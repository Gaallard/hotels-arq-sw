package users

import (
	dto "users-api/dto"
	models "users-api/model"

	e "users-api/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func RegisterUser(user models.User) (models.User, e.ApiError) {
	userExists := Db.Where("username = ?", user.Username).First(&user)
	if userExists.Error == nil {
		log.Error("El usuario ya existe")
		return user, e.NewBadRequestApiError("El usuario ya existe")
	}

	result := Db.Create(&user)
	log.Info("Usuario creado con exito")
	if result.Error != nil {
		log.Error("Error al crear el usuario")
		log.Error(result.Error)
		return user, e.NewBadRequestApiError("Error al crear usuario")
	}

	return user, nil
}

func GetUser(request dto.UserDto) (models.User, e.ApiError) {
	var user models.User

	result := Db.Where("username = ?", request.Username).First(&user)
	log.Println("result", result)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return user, e.NewBadRequestApiError("Error al buscar usuario")
	}

	return user, nil
}

func GetUserById(idUser int) (models.User, e.ApiError) {
	var userFound models.User
	result := Db.Where("id_user = ?", idUser).First(&userFound)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return userFound, e.NewBadRequestApiError("Error al buscar usuario")
	}

	return userFound, nil
}
