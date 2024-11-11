package usersController

import (
	Service "backend/services/users"
	"strings"

	"github.com/gin-gonic/gin"

	userDomain "backend/domain"

	middle "backend/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var userData userDomain.UserData
	c.BindJSON(&userData)

	loginResponse, err := Service.UserService.Login(userData)
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
func Extrac(c *gin.Context) {
	data := strings.TrimSpace(c.GetHeader("Authorization"))
	log.Println("token buscado: ", data)
	response, err := middle.ExtractClaims(data)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetUserByName(c *gin.Context) {

	var userDomain userDomain.UserData
	c.BindJSON(&userDomain)

	userDomain, err := Service.UserService.GetUserByName(userDomain)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDomain)

}

func UsuarioInsert(c *gin.Context) {
	var userDomain userDomain.UserData
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

	userDomain, er := Service.UserService.InsertUsuario(userDomain)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDomain)

}
