package main

import (
	"users-api/users-api/db"
	"users-api/users-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.StartDbEngine()
	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")
}
