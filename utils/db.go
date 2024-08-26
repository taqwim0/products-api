package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load("app.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}

	DB = db
	log.Println("Postgres database connection established")
}
