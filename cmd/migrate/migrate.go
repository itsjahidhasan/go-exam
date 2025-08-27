package main

import (
	"go-exam/config"
	"go-exam/db"
	"go-exam/db/migrate"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the migration direction: up or down")
	}
	dir := filepath.FromSlash("db/migrations")

	env := config.LoadConfig()
	conn, err := db.Open(env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	defer conn.Close()

	if err := migrate.Apply(conn, dir, os.Args[1]); err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Migration completed")
}
