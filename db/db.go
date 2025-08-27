package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Open(host, port, user, pass, name string) (*sql.DB, error) {
	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + name + " sslmode=disable"

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
