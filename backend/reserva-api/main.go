package main

import (
	"net/http"
	controller "reserva-api/controller"
	repo "reserva-api/repositories"
	service "reserva-api/services"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetReservaByID(ctx *gin.Context)
	updateReserva(ctx *gin.Context)
	InsertReserva(ctx *gin.Context)
}

func main() {
	sqlconfig := repo.SQLConfig{
		Name: "reservas",
		User: "root",
		Pass: "Tomas1927",
		Host: "localhost",
	}

	mainRepo := repo.NewSql(sqlconfig)
	Service := service.NewService(mainRepo)
	Controller := controller.NewController(Service)
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

	router.GET("/reservas/:id", Controller.GetReservaById)
	router.POST("/reservas/", Controller.InsertReserva)
	router.PUT("/reservas/", Controller.UpdateReserva)
	router.DELETE("/reservas/", Controller.DeleteReserva)
	router.Run(":8083")
}
