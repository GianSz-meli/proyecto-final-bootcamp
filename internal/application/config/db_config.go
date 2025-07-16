package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the MySQL database connection and returns the *sql.DB instance.
// The application will terminate if the connection cannot be established.
func InitDB() *sql.DB {
	db, err := ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mysql database connected")
	return db
}

// ConnectDB creates and opens a new MySQL database connection using environment variables.
// Returns the connection if successful, or an error if the connection fails.
func ConnectDB() (*sql.DB, error) {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
