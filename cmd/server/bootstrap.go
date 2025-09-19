package main

import (
	"fmt"

	"github.com/chinsiang99/simple-bank-go-project/internal/api/handlers"
	"github.com/chinsiang99/simple-bank-go-project/internal/api/routers"
	"github.com/chinsiang99/simple-bank-go-project/internal/config"
	"github.com/chinsiang99/simple-bank-go-project/internal/database"
	db "github.com/chinsiang99/simple-bank-go-project/internal/database/sqlc"
	"github.com/chinsiang99/simple-bank-go-project/internal/services"
	"github.com/chinsiang99/simple-bank-go-project/internal/utils/logger"
)

func bootstrap() {
	cfg := config.New()

	// Initialize logger with file rotation
	logger.Init(cfg.LOG)

	// Initialize database
	// dbConn, err := database.Init(cfg.DB)
	// if err != nil {
	// 	logger.Fatal("Failed to initialize database:", err)
	// }

	// if err := dbConn.Migrate(); err != nil {
	// 	logger.Fatalf("failed to run migrations: %v", err)
	// }

	// Initialize the context
	// ctx := context.Background()

	dbConn, err := database.Init(cfg.DB)
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	store := db.NewStore(dbConn.DB)

	// repos := repositories.NewRepositoryManager(dbConn)
	services := services.NewServiceManager(store)
	handlers := handlers.NewHandlerManager(services)
	router := routers.NewRouterManager(handlers, cfg.APP, cfg.SECURITY)

	addr := fmt.Sprintf(":%v", cfg.APP.AppPort)
	logger.Infof("Starting server on port %v", addr)
	if err := router.Run(addr); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
