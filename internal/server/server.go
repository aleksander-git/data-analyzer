package server

import (
	"github.com/aleksander-git/data-analyzer/internal/controllers"
	middleware "github.com/aleksander-git/data-analyzer/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger(logger))
	router.Use(gin.RecoveryWithWriter(logger.Writer()))

	// Пример маршрута
	router.GET("/", controllers.HelloHandler)
	router.GET("/error", controllers.ErrorHandler)

	return router
}
