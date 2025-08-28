package main

import (
	"go-exam/db"
	"go-exam/db/migrate"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the migration direction: up or down")
	}

	conn, err := db.Open()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	defer conn.Close()

	if err := migrate.Apply(conn, os.Args[1]); err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Migration completed")
}
