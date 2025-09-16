package main

import _ "github.com/chinsiang99/simple-bank-go-project/internal/docs"

// _ "github.com/simple-go-project/docs" // Import swagger docs

// @title github.com/simple-go-project API
// @version 1.0
// @description This is a sample API server with authentication
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	bootstrap()
}
