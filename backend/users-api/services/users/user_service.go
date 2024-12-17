package services

import (
	Domain "backend/domain"
	Model "backend/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"
	"time"

	e "backend/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
)

type userServiceInterface interface {
	GetUserByName(user Model.User) (Model.User, e.ApiError)
	InsertUsuario(usuarioDomain Model.User) (Model.User, e.ApiError)
}

type Service struct {
	UserService  userServiceInterface
	cacheService userServiceInterface
}

func NewService(UserService userServiceInterface, cacheService userServiceInterface) Service {
	return Service{
		UserService:  UserService,
		cacheService: cacheService,
	}
}

func (s Service) GetContainerStatus(containerName string) string {
	cmd := exec.Command("docker", "inspect", "--format", "{{.State.Status}}", containerName)
	output, err := cmd.Output()
	if err != nil {
		log.Errorf("Error getting status for container %s: %v", containerName, err)
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func (s Service) ManageContainer(containerName, action string) error {
	// Validar acción
	if action != "start" && action != "stop" {
		return fmt.Errorf("invalid action: %s", action)
	}

	cmd := exec.Command("docker", action, containerName)
	if err := cmd.Run(); err != nil {
		log.Errorf("Error managing container %s: %v", containerName, err)
		return fmt.Errorf("failed to %s container %s", action, containerName)
	}
	return nil
}

func (s Service) ListContainersStatus(containerNames []string) []Domain.ContainerStatus {
	var statuses []Domain.ContainerStatus
	for _, container := range containerNames {
		status := s.GetContainerStatus(container)
		statuses = append(statuses, Domain.ContainerStatus{Name: container, Status: status})
	}
	return statuses
}

func (s Service) GetUserByName(usuarioDomain Domain.UserData) (Domain.UserData, e.ApiError) {

	usuario := Model.User{
		User: usuarioDomain.User,
	}

	user, err := s.cacheService.GetUserByName(usuario)
	if err != nil {
		user, err = s.UserService.GetUserByName(usuario)

		if err != nil {
			return Domain.UserData{}, e.NewBadRequestApiError("Error al buscar el usuario")
		}

		if _, err := s.cacheService.InsertUsuario(usuario); err != nil {
			return Domain.UserData{}, e.NewBadRequestApiError("Error al insertar el usuario a cache")
		}
	}
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

func (s Service) InsertUsuario(usuarioDomain Domain.UserData) (Domain.UserData, e.ApiError) {

	usuario := Model.User{
		User:  usuarioDomain.User,
		Admin: usuarioDomain.Admin,
	}

	var result, er = s.UserService.GetUserByName(usuario)

	if er != nil {

		hash := md5.New()
		hash.Write([]byte(usuarioDomain.Password))
		usuario.Password = hex.EncodeToString(hash.Sum(nil))

		usuario2, err := s.UserService.InsertUsuario(usuario)

		if err != nil {
			return usuarioDomain, e.NewBadRequestApiError("usuario no inseertado")
		}

		_, err = s.cacheService.InsertUsuario(usuario2)

		if err != nil {
			return usuarioDomain, e.NewBadRequestApiError("usuario no insertado en cache")
		}

		usuarioDomain.Id = usuario2.Id

		return usuarioDomain, nil
	}

	return Domain.UserData(result), e.NewBadRequestApiError("Nombre de usuario existente")

}

func (s Service) Login(User Domain.UserData) (Domain.LoginData, e.ApiError) {

	usuario := Model.User{
		User:  User.User,
		Admin: User.Admin,
	}

	var user, err = s.cacheService.GetUserByName(usuario)
	if err != nil {
		user, err = s.cacheService.GetUserByName(usuario)

	}
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
