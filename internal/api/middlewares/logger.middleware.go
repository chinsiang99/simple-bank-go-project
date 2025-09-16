package middlewares

import (
	"time"

	"github.com/chinsiang99/simple-bank-go-project/internal/utils/logger"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Skip noisy OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// Process request
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		latencyMs := float64(latency.Microseconds()) / 1000.0 // ms

		fields := map[string]interface{}{
			"client_ip": c.ClientIP(),
			"status":    c.Writer.Status(),
			"latency":   latencyMs,
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"query":     c.Request.URL.RawQuery,
			"type":      "http_request",
			"timestamp": end.Format("2006/01/02 15:04:05"),
		}

		if c.Writer.Status() >= 400 {
			fields["errors"] = c.Errors.ByType(gin.ErrorTypePrivate).String()
			logger.WithFields(fields).Error("GIN Request")
		} else {
			logger.WithFields(fields).Info("GIN Request")
		}
	}
}
