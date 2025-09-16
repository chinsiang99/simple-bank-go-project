package middlewares

// import "github.com/gin-gonic/gin"

// // OWASPHeaders sets basic HTTP headers to improve security
// func OWASPHeaders() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Prevent clickjacking by not allowing this page to be framed
// 		c.Writer.Header().Set("X-Frame-Options", "DENY")

// 		// Prevent MIME type sniffing, reduces risk of XSS attacks
// 		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

// 		// Enable browser XSS protection (basic level)
// 		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

// 		// Restrict information sent in the Referer header
// 		c.Writer.Header().Set("Referrer-Policy", "no-referrer")

// 		// Restrict sources from which browser can load content
// 		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")

// 		// Enforce HTTPS connections for future requests
// 		c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

// 		c.Next() // Continue to next middleware or route handler
// 	}
// }
