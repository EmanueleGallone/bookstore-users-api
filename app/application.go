package app

import (
	"bookstore-users-api/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

var ( //router will be available only in application package
	router            = gin.Default()
	port, portPresent = os.LookupEnv("PORT") //getting port from dockerfile
)

func StartApplication() {
	mapURLS()

	logger.Info("Starting Application")

	if !portPresent {
		port = strconv.Itoa(8080)
	}

	err := router.Run(fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Error while running the router: %v", err)
	}
}
