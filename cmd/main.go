package main

import (
	"os"

	"github.com/aleksander-git/data-analyzer/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.InfoLevel)

	logger.Info("Logger initialized")
	logger.Info("Starting server")

	router := server.NewRouter(logger)
	if err := router.Run(":8080"); err != nil {
		logger.Fatalf("could not start server: %v", err)
	}
}
