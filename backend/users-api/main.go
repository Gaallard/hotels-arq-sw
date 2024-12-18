package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	clients "backend/clients/users"
	controller "backend/controllers/users"
	service "backend/services/users"

	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, X-Auth-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	cacheConfig := clients.CacheConfig{
		MaxSize:      100000,
		ItemsToPrune: 100,
		Duration:     5 * time.Minute,
	}

	sqlconfig := clients.MySQLConfig{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		//Name: "users",
		//User: "root",
		//Pass: "root",
		//Host: os.Getenv("DB_HOST"),
		//Name: "users",
		//	User: "root",
		//Pass: "root",
		//Host: "mysql-container",
	}

	mainRepo := clients.NewSql(sqlconfig)
	cacheRepo := clients.NewCache(cacheConfig)

	service := service.NewService(mainRepo, cacheRepo)
	controller := controller.NewController(service)

	router.POST("/users", controller.UsuarioInsert)
	router.POST("/users/login", controller.Login)
	router.GET("/users/token", controller.Extrac)
	router.GET("/users/cache", controller.GetUserByName)

	router.GET("/users", controller.GetUserByName)
	router.GET("/users/containers", controller.GetContainers)
	router.POST("/users/containers/:action/:name", controller.ManageContainer)

	router.Run(":8080")

}
