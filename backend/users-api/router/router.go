package router

import (
	"net/http"
	"time"
	users "users-api/controller"
	"users-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetHotelByIDFromHotelsAPI(c *gin.Context) {
	hotelID := c.Param("_id")
	hotelsAPIURL := "http://hotels-api:8080/hotels/" + hotelID

	resp, err := http.Get(hotelsAPIURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar con hotels-api"})
		return
	}
	defer resp.Body.Close()

	// Pasar la respuesta de hotels-api al cliente
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

func MapUrls(engine *gin.Engine) {
	engine.Use(middleware.CORSMiddleware())
	engine.POST("/users/login", users.Login)
	engine.POST("/users/register", users.RegisterUser)
	engine.GET("users/:id", users.GetUserById)
	authorized := engine.Group("")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/hotels/:_id", GetHotelByIDFromHotelsAPI)

		admin := authorized.Group("")
		admin.Use(middleware.AuthMiddlewareAdmin())
		/*
			{
				admin.POST("/hotels-api/hotels", hotels.InsertHotel)
				admin.PUT("/hotels-api/hotels/:_id", hotels.UpdateHotel)
			}*/
	}
}

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
