package database

import (
	"fmt"
	"golang-todo-api/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:root@tcp(46.250.236.154:3391)/todo-api-golang?charset=utf8mb4&parseTime=True&loc=Local"

func Init() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer DisconnectDB(db)

	// Migrate the schema
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Seed the database
	_Seeder(db)

	log.Println("Database initialized with migration and seeding.")
}

// Connect to the database
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Configure connection pooling
	mysqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL DB instance: %w", err)
	}
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to the database successfully.")
	return db, nil
}

func DisconnectDB(db *gorm.DB) error {
	mysqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL DB instance: %w", err)
	}
	if err := mysqlDB.Close(); err != nil {
		return fmt.Errorf("error closing the database connection: %w", err)
	}
	log.Println("Disconnected from the database successfully.")
	return nil
}
