package main

import (
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

}
