package database

import (
	"fmt"
	"log"
	"time"

	// "trouble-ticket-ms/src/models"

	"github.com/chinsiang99/simple-bank-go-project/internal/config"
	"github.com/chinsiang99/simple-bank-go-project/internal/models"
	"github.com/chinsiang99/simple-bank-go-project/internal/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type Config struct {
	Host       string
	Name       string
	User       string
	Pass       string
	Port       string
	Schema     string
	MaxOpenCon int
	MaxIdleCon int
}

func Init(config *config.DBConfig) (*DB, error) {
	// config := config.New()

	// MySQL DSN format
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
	// 	config.User,
	// 	config.Pass,
	// 	config.Host,
	// 	config.Port,
	// 	config.Name,
	// )

	// postgres dsn format
	// note that only sslmode=disable when it is been used in local, remove it for production
	// below commented code is for local or docker desktop use only
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Pass,
		config.Name,
		config.Schema,
	)

	// below code is for production use that has ssl enabled
	// dsn := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s search_path=%s",
	// 	config.Host,
	// 	config.Port,
	// 	config.User,
	// 	config.Pass,
	// 	config.Name,
	// 	config.Schema,
	// )

	var dbConn *gorm.DB
	var err error

	// Retry up to 10 times to connect to MySQL, as the service might not be ready immediately after container startup (Docker Compose *)
	for i := 0; i < 10; i++ {
		dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err == nil {
			break
		}
		logger.Debugf("Failed to open DB connection (attempt %d): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		logger.Errorf("Failed to open DB connection after retries: %v", err)
		log.Panicf("Failed to open DB connection after retries: %v", err)
	}

	// Get the underlying *sql.DB object
	sqlDB, err := dbConn.DB()

	if err != nil {
		logger.Error(err.Error())
		log.Panic(err.Error())
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenCon)
	sqlDB.SetMaxIdleConns(config.MaxIdleCon)

	log.Printf("Successfully connected to %s database on %s:%s", config.Name, config.Host, config.Port)

	return &DB{dbConn}, nil
}

// func (db *DB) MigrationUpToDate() bool {
// 	// migrator := db.DB.Migrator()
// 	// if !migrator.HasTable(&models.TroubleTicket{}) {
// 	// 	logger.Debug("table does not exist")
// 	// 	log.Println("table does not exist")
// 	// 	return false
// 	// }
// 	// return true
// }

func (db *DB) Migrate() error {
	return db.AutoMigrate(
		&models.User{},
		// &models.Ticket{},  // add more models here
	)
}

/*
	func CloseDB(db *DB) {
		sqlDB, _ := DB()
		err := sqlClose()
		if err != nil {
			panic(err)
		}
		log.Printf("DB Closed Successfully")
	}
*/
