package main

import (
	"log"
	clients "search-api/clients/queues"
	controllers "search-api/controller"
	search "search-api/respositories"
	services "search-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Solr
	solrConfig := search.SolrConfig{
		BaseURL:    "",
		Collection: "hotels",
	}

	rabbit := clients.RabbitConfig{
		Host:      "localhost",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "hoteUCC",
	}

	rabbitRpo := clients.NewRabbit(rabbit)
	solrRepo := search.NewSolr(solrConfig)

	// Services, le di el rabbit al service
	service := services.NewService(solrRepo, rabbitRpo)

	// Handlers
	controller := controllers.NewController(service)

	// Create router
	router := gin.Default()

	// URL mappings
	// /hotels/search?q=sheraton&limit=10&offset=0
	router.GET("/hotels/search", controller.Search)

	// Run application
	if err := router.Run(":8081"); err != nil {
		log.Panicf("Error running application: %v", err)
	}
}
