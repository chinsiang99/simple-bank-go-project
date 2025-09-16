package middlewares

// import "github.com/gin-gonic/gin"

// // CORS middleware sets headers to allow cross-origin requests
// func CORS() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Short-circuit OPTIONS requests (preflight) to avoid unnecessary processing
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204) // 204 No Content
// 			return
// 		}

// 		// Allow specific HTTP methods
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

// 		// Allow all origins to access API
// 		// Adjust "*" to specific domains in production for better security
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

// 		// Allow specific headers from client
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

// 		// Headers exposed to client
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

// 		// Allow credentials (cookies, authorization headers) to be sent in cross-origin requests
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		c.Next() // Continue processing actual request
// 	}
// }
