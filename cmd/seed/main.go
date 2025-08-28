package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-exam/db"
	"log"
	"os"
)

type seedUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func seedUsers(tx *sql.Tx) error {
	b, err := os.ReadFile("seed/users.json")
	if err != nil {
		return err
	}

	var users []seedUser
	if err := json.Unmarshal(b, &users); err != nil {
		return err
	}

	ctx := context.Background()
	for _, u := range users {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
			u.Name, u.Email, u.Age)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	conn, err := db.Open()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal("Error starting transaction: ", err)
	}
	defer tx.Rollback()

	if err := seedUsers(tx); err != nil {
		log.Fatal("Error seeding users: ", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	log.Println("✔ Seeding done")
}
