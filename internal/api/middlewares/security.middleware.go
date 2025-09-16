package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/chinsiang99/simple-bank-go-project/internal/config"
	"github.com/gin-gonic/gin"
)

// SecurityMiddleware handles:
// 1. Preflight requests (OPTIONS)
// 2. CORS
// 3. OWASP security headers (production only)
func SecurityMiddleware(appConfig *config.AppConfig, securityConfig *config.SecurityConfig) gin.HandlerFunc {
	// Read CORS config from environment
	allowOrigins := strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	allowMethods := os.Getenv("CORS_ALLOW_METHODS")
	allowHeaders := os.Getenv("CORS_ALLOW_HEADERS")
	exposeHeaders := os.Getenv("CORS_EXPOSE_HEADERS")
	allowCredentials := os.Getenv("CORS_ALLOW_CREDENTIALS") == "true"
	maxAge := os.Getenv("CORS_MAX_AGE")

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		// ---------------------------
		// Step 1: Handle Preflight
		// ---------------------------
		if c.Request.Method == http.MethodOptions {
			if isOriginAllowed(origin, allowOrigins) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				c.Writer.Header().Set("Vary", "Origin")
				c.Writer.Header().Set("Access-Control-Allow-Methods", allowMethods)
				c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
				c.Writer.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
				c.Writer.Header().Set("Access-Control-Max-Age", maxAge)
				if allowCredentials {
					c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				}
			}
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// ---------------------------
		// Step 2: Apply CORS Headers
		// ---------------------------
		if isOriginAllowed(origin, allowOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
			c.Writer.Header().Set("Access-Control-Allow-Methods", allowMethods)
			c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
			c.Writer.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
			if allowCredentials {
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		// ---------------------------
		// Step 3: OWASP Security Headers (production only)
		// ---------------------------
		if appConfig.Environment == "production" {
			c.Writer.Header().Set("X-Frame-Options", "DENY")           // Prevent clickjacking
			c.Writer.Header().Set("X-Content-Type-Options", "nosniff") // Prevent MIME sniffing
			c.Writer.Header().Set("X-XSS-Protection", "1; mode=block") // Enable XSS protection
			c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self'; frame-ancestors 'none';")
			c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload") // HSTS
			c.Writer.Header().Set("Referrer-Policy", "no-referrer")                                            // Hide referrer info
		}

		// Continue to next handler
		c.Next()
	}
}

// Helper: check if the request origin is allowed
func isOriginAllowed(origin string, allowOrigins []string) bool {
	for _, o := range allowOrigins {
		if strings.TrimSpace(o) == origin {
			return true
		}
	}
	return false
}
