package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getConnectionString() string {
	loadEnv()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}

func ConnectDatabase() {
	connectionString := getConnectionString()
	var err error
	DB, err = sql.Open(("postgres"), connectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping to database:", err)
	}
}
