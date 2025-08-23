package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Get the database source URL from the environment variables we set in docker-compose
	dsn := os.Getenv("DB_SOURCE")
	if dsn == "" {
		log.Fatal("DB_SOURCE environment variable is not set")
	}

	//open a connection on the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	err = database.AutoMigrate(&Experiment{}, &Variation{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	DB = database
	log.Println("Database connection succesfull and migrations are up-to-date.")
}