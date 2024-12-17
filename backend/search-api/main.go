package main

import (
	"log"
	"net/http"
	"search-api/clients/queues"
	controllers "search-api/controller"
	repositories "search-api/respositories"
	services "search-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Solr
	solrRepo := repositories.NewSolr(repositories.SolrConfig{
		Host:       "solr",   // Solr host
		Port:       "8983",   // Solr port
		Collection: "hotels", // Collection name
	})

	// Rabbit
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "rabbit",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "hotelUCC",
	})

	// Hotels API
	hotelsAPI := repositories.NewHTTP(repositories.HTTPConfig{
		Host: "nginx",
		Port: "8081",
	})

	// Services
	service := services.NewService(solrRepo, hotelsAPI)

	// Controllers
	controller := controllers.NewController(service)

	// Launch rabbit consumer
	if err := eventsQueue.StartConsumer(service.HandleHotelNew); err != nil {
		log.Fatalf("Error running consumer: %v", err)
	}

	// Create router
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

	router.GET("/search", controller.Search)
	if err := router.Run(":8084"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
