package main

import (
	hotelsController "hotels-api/controllers/hotels"
	hotelsRepository "hotels-api/repositories/hotels"
	hotelsService "hotels-api/services/hotels"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Config
	cacheConfig := hotelsRepository.CacheConfig{
		MaxSize:      100000,
		ItemsToPrune: 100,
	}

	mongoConfig := hotelsRepository.MongoConfig{
		Host:       "localhost",
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "hotels",
		Collection: "hotels",
	}

	// Dependencies
	mainRepository := hotelsRepository.NewMongo(mongoConfig)
	cacheRepository := hotelsRepository.NewCache(cacheConfig)
	service := hotelsService.NewService(mainRepository, cacheRepository)
	controller := hotelsController.NewController(service)

	// Router
	router.GET("/hotels/:_id", controller.GetHotelByID)
	router.POST("/hotels", controller.InsertHotel)
	router.PUT("hotels/:_id", controller.UpdateHotel)

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
