package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetDSN() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Tidak dapat memuat file .env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Debug logging to check environment variables
	log.Printf("Database configuration: host=%s, port=%s, user=%s, dbname=%s", host, port, user, dbname)

	encodedPassword := url.QueryEscape(password)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, encodedPassword, host, port, dbname)
	log.Printf("DSN: %s", dsn)

	return dsn
}

func ConnectDB() (*sql.DB, error) {

	dsn := GetDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("berhasil tersambung ke database")
	return db, nil
}
