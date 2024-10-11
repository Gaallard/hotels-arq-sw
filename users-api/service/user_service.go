package users

import (
	"crypto/md5"
	"encoding/hex"
	client "users-api/users-api/client"
	dto "users-api/users-api/dto"
	token "users-api/users-api/dto"
	e "users-api/users-api/errors"
	model "users-api/users-api/model"

	"github.com/dgrijalva/jwt-go"
)

func RegisterUser(request dto.UserDto) (dto.UserDto, e.ApiError) {
	var user model.User

	user.Username = request.Username
	user.Password = request.Password

	hash := md5.New()
	hash.Write([]byte(request.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	user, err := client.RegisterUser(user)
	if err != nil {
		return request, e.NewBadRequestApiError("Error al registrar usuario")
	}
	request.ID = user.Id_user

	return request, nil
}

var jwtKey = []byte("secret_key")

func Login(request dto.UserDto) (token.TokenDto, e.ApiError) {
	var user, err = client.GetUser(request)
	var tokenDto token.TokenDto
	if err != nil {
		return tokenDto, e.NewBadRequestApiError("Usuario no encontrado")
	}
	var pswMd5 = md5.Sum([]byte(request.Password))
	pswMd5String := hex.EncodeToString(pswMd5[:])

	if pswMd5String == user.Password {
		tokn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id_user,
		})
		tokenString, _ := tokn.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.Id_user = user.Id_user
		tokenDto.Role = user.Role
		return tokenDto, nil
	} else {
		return tokenDto, e.NewBadRequestApiError("contrasenia incorrecta")
	}

}
