package main

import (
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
		Pass: "Tomihuspenina2003",
		Host: "localhost",
	}

	mainRepo := repo.NewSql(sqlconfig)
	Service := service.NewService(mainRepo)
	Controller := controller.NewController(Service)
	router := gin.Default()
	router.GET("/reservas/:id", Controller.GetReservaById)
	router.POST("/reservas/", Controller.InsertReserva)
	router.PUT("/reservas/", Controller.UpdateReserva)
	router.DELETE("/reservas/", Controller.DeleteReserva)
	router.Run(":8083")
}
