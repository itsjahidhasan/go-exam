package main

import (
	"fmt"
	"go-exam/config"
	"go-exam/db"
	"go-exam/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Go Exam Application")
	// Application entry point
	env := config.LoadConfig()
	conn, err := db.Open(env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	defer conn.Close()

	router := http.NewServeMux()
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Go Exam Application"))
	})

	// here will be routes
	routes.UserRoutes(router)

	fmt.Println("Server: " + "http://localhost:" + env.AppPort + "/")
	http.ListenAndServe(":"+env.AppPort, router)

}
