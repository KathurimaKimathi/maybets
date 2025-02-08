package gorm

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DBInstance holds the database connection
type DBInstance struct {
	DB *gorm.DB
}

// NewDBInstance initializes a new SQLite database instance
func NewDBInstance() (*DBInstance, error) {
	db, err := startDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to start database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(80)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &DBInstance{DB: db}, nil
}

// startDatabase initializes the SQLite database
func startDatabase() (*gorm.DB, error) {
	// Use a persistent SQLite file instead of an in-memory database
	path := os.Getenv("SQLITE_URL")
	if path == "" {
		path = "../../../../../.."
	}

	dbPath := fmt.Sprintf("%s/bets.db", path)

	// Ensure the file exists
	file, err := os.Create(dbPath)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create SQLite DB file: %v", err)
	}

	file.Close()

	return boot(dbPath)
}

// boot initializes GORM with SQLite
func boot(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		CreateBatchSize: 1000,
		Logger:          logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Errorf("Failed to connect to database: %s", err)
		return nil, err
	}

	// Check connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("Failed to get DB instance: %s", err)
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Errorf("Unable to ping the database: %s", err)
		return nil, err
	}

	// Add OpenTelemetry plugin for tracing
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Errorf("Unable to add otel plugin: %s", err)
		return nil, err
	}

	return db, nil
}
