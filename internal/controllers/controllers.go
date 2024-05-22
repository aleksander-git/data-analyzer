package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func ErrorHandler(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
	c.Error(errors.New("this is a test error"))
}
