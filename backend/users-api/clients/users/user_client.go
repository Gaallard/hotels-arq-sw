package clientUsers

import (
	Domain "backend/domain"
	e "backend/errors"
	Model "backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserByName(Usuario Domain.UserData) (Model.User, e.ApiError) {
	var user Model.User

	result := Db.Where("User = ?", Usuario.User).First(&user)
	log.Debug("User: ", user)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return user, e.NewBadRequestApiError("Error al buscar el usuario")
	}

	return user, nil
}

func InsertUsuario(user Model.User) (Model.User, e.ApiError) {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("Error al crear el usuario")
		log.Error(result.Error)
		return user, e.NewBadRequestApiError("Error al crear el usuario")
	}
	log.Debug("User Created: ", user.Id)
	return user, nil
}

func GetUserById(Id int) (Model.User, e.ApiError) {
	var userId Model.User

	result := Db.Where("id = ?", Id).First(&userId)
	log.Debug("id: ", userId)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return userId, e.NewBadRequestApiError("Error al buscar el usuario")
	}

	return userId, nil
}

func GetuserName(buscado int) (string, e.ApiError) {
	var userId Model.User

	result := Db.Where("id = ?", buscado).First(&userId)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return userId.User, e.NewBadRequestApiError("Error al buscar el usuario")
	}

	return userId.User, nil
}
