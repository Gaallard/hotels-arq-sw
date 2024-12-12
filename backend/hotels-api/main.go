package main

import (
	"context"
	hotelsController "hotels-api/controllers/hotels"
	hotelsRepository "hotels-api/repositories/hotels"
	hotelsService "hotels-api/services/hotels"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// Config
	cacheConfig := hotelsRepository.CacheConfig{
		MaxSize:      100000,
		ItemsToPrune: 100,
	}

	//username y password de tomi: root / root
	mongoConfig := hotelsRepository.MongoConfig{
		Host:       "mongo",
		Port:       "27017",
		Username:   "root", //fran:
		Password:   "root", //fran:
		Database:   "hotels",
		Collection: "hotels",
	}

	HostR := "rabbit"
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
	router.GET("/hotels/:id", controller.GetHotelByID)
	router.POST("/hotels", controller.InsertHotel)
	router.GET("/hotels/available-rooms/:id", controller.GetAvailableRooms)
	router.GET("/hotels", controller.GetAllHotels2)

	router.PUT("/hotels/:id", controller.UpdateHotel)

	log.Println("Servidor corriendo en http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
