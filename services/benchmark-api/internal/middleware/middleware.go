package middleware

import (
	"benchmark-api/pkg/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(_ gin.LogFormatterParams) string {
		return ""
	})
}

func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Request-ID", generateRequestID())
		c.Next()
	}
}

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := string(rune(c.Writer.Status()))

		metrics.RecordHTTPRequest(c.Request.Method, c.FullPath(), status, duration)
	}
}

func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + "req"
}
