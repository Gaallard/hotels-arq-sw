package main

import (
	"context"
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
		Username:   "frantmateos", //fran:
		Password:   "Tomas1927",   //fran:
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

	ctx := context.Background()

	err := service.GetAllHotels(ctx)
	if err != nil {
		log.Fatalf("error publishing hotels to RabbitMQ on startup: %v", err)
	}

	// Router
	router.GET("/hotels/:_id", controller.GetHotelByID)
	router.POST("/hotels", controller.InsertHotel)
	router.PUT("hotels/:_id", controller.UpdateHotel)
	router.GET("/hotels/available-rooms", controller.GetAvailableRooms)

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
