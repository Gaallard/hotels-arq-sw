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

	//username y password de tomi: root / root
	mongoConfig := hotelsRepository.MongoConfig{
		Host:       "localhost",
		Port:       "27017",
		Username:   "root", //fran:
		Password:   "root", //fran:
		Database:   "hotels",
		Collection: "hotels",
	}

	HostR := "localhost"
	PortR := "5672"
	UsernameR := "guest"
	PasswordR := "guest"
	QueueR := "hotelUCC"

	// Dependencies
	rabbitRpo := hotelsRepository.NewPublisher(UsernameR, PasswordR, HostR, PortR, QueueR)
	mainRepository := hotelsRepository.NewMongo(mongoConfig)
	cacheRepository := hotelsRepository.NewCache(cacheConfig)
	//le el rabiitpubliser al service
	service := hotelsService.NewService(mainRepository, cacheRepository, rabbitRpo)
	controller := hotelsController.NewController(service)

	// Router
	router.GET("/hotels/:id", controller.GetHotelByID)
	router.POST("/hotels", controller.InsertHotel)
	router.PUT("/hotels/:id", controller.UpdateHotel)

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
