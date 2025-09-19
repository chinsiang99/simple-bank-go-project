package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/chinsiang99/simple-bank-go-project/internal/config"
	"github.com/chinsiang99/simple-bank-go-project/internal/utils/logger"

	_ "github.com/lib/pq" // Postgres driver
)

type DB struct {
	*sql.DB
}

// Init opens a plain sql.DB connection (for use with sqlc)
func Init(cfg *config.DBConfig) (*DB, error) {
	// DSN format for Postgres
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name, cfg.Schema,
	)

	var sqlDB *sql.DB
	var err error

	// Retry up to 10 times in case DB is not ready yet
	for i := 0; i < 10; i++ {
		sqlDB, err = sql.Open("postgres", dsn)
		if err == nil {
			// Ping to confirm connection
			err = sqlDB.Ping()
			if err == nil {
				break
			}
		}
		logger.Debugf("Failed to open DB connection (attempt %d): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		logger.Errorf("Failed to open DB connection after retries: %v", err)
		log.Panicf("Failed to open DB connection after retries: %v", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenCon)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleCon)

	log.Printf("✅ Successfully connected to %s database on %s:%s", cfg.Name, cfg.Host, cfg.Port)

	return &DB{sqlDB}, nil
}
