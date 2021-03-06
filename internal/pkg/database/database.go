package database

import (
	"fmt"
	"github.com/devgit072/production-ready-rest-api-in-go/internal/books"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// Returns a pointer to databse object.
func NewDatabase() (*gorm.DB, error) {
	log.Println("Setting up new DB connection")
	// We should not store the credentials in the file because we will be commiting the code in the github.
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUsername, dbPassword, dbName, "disable")
	db, err := gorm.Open(postgres.Open(str))
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return nil, err
	}
	d, err := db.DB()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return nil, err
	}
	if err := d.Ping(); err != nil {
		log.Fatalf("Error: %s", err.Error())
		return nil, err
	}
	log.Println("New database connection created successfully")
	return db, nil
}

func AutoMigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&books.Book{}); err != nil {
		return err
	}
	return nil
}
