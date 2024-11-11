package services

import (
	Domain "backend/domain"
	Model "backend/model"
	"crypto/md5"
	"encoding/hex"
	"time"

	Clients "backend/clients/users"

	e "backend/errors"

	jwt "github.com/dgrijalva/jwt-go"
)

type userService struct{}

type userServiceInterface interface {
	GetUserByName(user Domain.UserData) (Domain.UserData, e.ApiError)
	InsertUsuario(usuarioDomain Domain.UserData) (Domain.UserData, e.ApiError)
	Login(User Domain.UserData) (Domain.LoginData, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserByName(usuario Domain.UserData) (Domain.UserData, e.ApiError) {

	var user, err = Clients.GetUserByName(usuario)
	var userDomain Domain.UserData

	if err != nil {
		return userDomain, e.NewBadRequestApiError("usuario no encontrado")
	}

	userDomain.Id = user.Id
	userDomain.User = user.User
	userDomain.Password = user.Password
	userDomain.Admin = user.Admin

	return userDomain, nil

}

func (u *userService) InsertUsuario(usuarioDomain Domain.UserData) (Domain.UserData, e.ApiError) {

	var usuario Model.User
	var result, er = Clients.GetUserByName(usuarioDomain)

	if er != nil {

		usuario.User = usuarioDomain.User
		usuario.Password = usuarioDomain.Password
		usuario.Admin = usuarioDomain.Admin

		hash := md5.New()
		hash.Write([]byte(usuarioDomain.Password))
		usuario.Password = hex.EncodeToString(hash.Sum(nil))

		var usuario2, err = Clients.InsertUsuario(usuario)

		if err != nil {
			return usuarioDomain, e.NewBadRequestApiError("usuario no inseertado")
		}

		usuarioDomain.Id = usuario2.Id
		usuarioDomain.Admin = usuario.Admin

		return usuarioDomain, nil
	}

	return Domain.UserData(result), e.NewBadRequestApiError("Nombre de usuario existente")

}

func (u *userService) Login(User Domain.UserData) (Domain.LoginData, e.ApiError) {
	var user, err = Clients.GetUserByName(User)
	var tokenDomain Domain.LoginData

	if err != nil {
		return tokenDomain, e.NewBadRequestApiError("usuario no encontrado")
	}

	var Logpsw = md5.Sum([]byte(User.Password))
	psw := hex.EncodeToString(Logpsw[:])

	if psw == user.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"idU":    user.Id,
			"Adminu": user.Admin,
			"exp":    time.Now().Add(time.Hour * 72).Unix(),
		})
		t, _ := token.SignedString([]byte("frantomi"))
		tokenDomain.Token = t
		tokenDomain.IdU = user.Id
		tokenDomain.AdminU = user.Admin
		return tokenDomain, nil
	} else {
		return tokenDomain, e.NewBadRequestApiError("Contrasenia incorrecta")
	}

}
