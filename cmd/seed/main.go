package main

import (
	"context"
	"encoding/json"
	"go-exam/config"
	"go-exam/db"
	"log"
	"os"
)

type seedUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	env := config.LoadConfig()
	conn, err := db.Open(env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	defer conn.Close()

	b, err := os.ReadFile("seed.json")
	if err != nil {
		log.Fatal("Error reading seed file: ", err)
	}

	var users []seedUser
	if err := json.Unmarshal(b, &users); err != nil {
		log.Fatal("Error unmarshalling seed data: ", err)
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal("Error starting transaction: ", err)
	}
	defer tx.Rollback()
	ctx := context.Background()
	for _, u := range users {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", u.Name, u.Email, u.Age)
		if err != nil {
			log.Fatal("Error inserting user: ", err)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
	log.Println("✔ seeding done")

}
