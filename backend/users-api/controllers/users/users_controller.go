package usersController

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	Domain "backend/domain"

	e "backend/errors"
	middle "backend/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type UserService interface {
	GetUserByName(usuarioDomain Domain.UserData) (Domain.UserData, e.ApiError)
	Login(User Domain.UserData) (Domain.LoginData, e.ApiError)
	InsertUsuario(usuarioDomain Domain.UserData) (Domain.UserData, e.ApiError)
	GetContainerStatus(containerName string) string
	GetContainersStatus(containerNames []string) []Domain.ContainerStatus
	ManageContainer(containerName, action string) error
}

type Controller struct {
	service UserService
}

func NewController(service UserService) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) Login(c *gin.Context) {
	var userData Domain.UserData
	c.BindJSON(&userData)

	loginResponse, err := controller.service.Login(userData)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	/*
		cookie := &http.Cookie{
			Name:     "A1uthorization",
			Value:    loginResponse.Token,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(c.Writer, cookie)
		//c.SetCookie("1Authorization", loginResponse.Token, 3600, "/", "localhost", true, false)
	*/
	c.JSON(http.StatusOK, loginResponse)
}
func (controller Controller) Extrac(c *gin.Context) {
	data := strings.TrimSpace(c.GetHeader("Authorization"))
	log.Println("token buscado: ", data)
	response, err := middle.ExtractClaims(data)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller Controller) GetUserByName(c *gin.Context) {

	var userDomain Domain.UserData
	c.BindJSON(&userDomain)

	userDomain, err := controller.service.GetUserByName(userDomain)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDomain)

}

func (controller Controller) UsuarioInsert(c *gin.Context) {
	var userDomain Domain.UserData
	err := c.BindJSON(&userDomain)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if userDomain.Admin {
		log.Info("creating admin user")
	} else {
		log.Info("creating regular user")
	}

	userDomain, er := controller.service.InsertUsuario(userDomain)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDomain)

}
func (controller Controller) GetContainers(c *gin.Context) {
	containers := []string{
		"hotels-api-container-1", "hotels-api-container-2", "hotels-api-container-3",
	}

	statuses := controller.service.GetContainersStatus(containers)
	c.JSON(http.StatusOK, statuses)
}

func (controller Controller) ManageContainer(c *gin.Context) {
	action := c.Param("action")
	containerName := c.Param("name")

	err := controller.service.ManageContainer(containerName, action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Salida: ": fmt.Sprintf("Container %s %s correctamente", containerName, action)})
}
