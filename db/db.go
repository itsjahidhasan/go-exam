package db

import (
	"database/sql"
	"go-exam/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Open() (*sql.DB, error) {
	env := config.LoadConfig()
	dsn := "host=" + env.DBHost + " port=" + env.DBPort + " user=" + env.DBUser + " password=" + env.DBPass + " dbname=" + env.DBName + " sslmode=disable"

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := dbConn.Ping(); err != nil {
		return nil, err
	}

	DB = dbConn // assign to global variable so other packages can use it
	log.Println("Connected to PostgreSQL successfully")

	return DB, nil
}
