package app

import (
	"bookstore-users-api/logger"
	"github.com/gin-gonic/gin"
	"log"
)

var ( //router will be available only in application package
	router = gin.Default()
)

func StartApplication() {
	mapURLS()

	logger.Info("Starting Application")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error while running the router: %v", err)
	}
}
