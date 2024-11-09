package main

import (
	"log"
	"search-api/clients/queues"
	controllers "search-api/controller"
	repositories "search-api/respositories"
	services "search-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Solr
	solrRepo := repositories.NewSolr(repositories.SolrConfig{
		Host:       "localhost", // Solr host
		Port:       "8983",      // Solr port
		Collection: "hotels",    // Collection name
	})

	// Rabbit
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "localhost",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "hotelUCC",
	})

	// Hotels API
	hotelsAPI := repositories.NewHTTP(repositories.HTTPConfig{
		Host: "localhost",
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
	router.GET("/search", controller.Search)
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
