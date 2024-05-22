package middleware

//требуется доработка. Нечитабельные логи

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logFields := logrus.Fields{
			"status":       status,
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"latency":      latency,
			"client_ip":    c.ClientIP(),
			"user_agent":   c.Request.UserAgent(),
			"referer":      c.Request.Referer(),
			"headers":      c.Request.Header,
			"request_body": requestBody,
		}

		if len(c.Errors) > 0 {
			errors := c.Errors.String()
			logger.WithFields(logFields).WithField("errors", errors).Error("Request details with errors")
		} else {
			logger.WithFields(logFields).Info("Request details")
		}
	}
}
