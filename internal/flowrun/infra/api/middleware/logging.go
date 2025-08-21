package middleware

import (
	"flow-run/internal/lib/logger"

	"github.com/gin-gonic/gin"
)

type LoggingMiddleware struct{}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Log.Infof("Request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
		logger.Log.Infof("Response: %d", c.Writer.Status())
	}
}
