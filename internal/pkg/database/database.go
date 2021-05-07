package database

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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
