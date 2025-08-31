package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-exam/db"
	"go-exam/modals"
	"log"
	"os"
)

type seedUser struct {
	Name        string `json:"name"`
	CashBalance string `json:"cashBalance"`
}
type seedRestaurant struct {
	Name        string `json:"name"`
	CashBalance string `json:"cashBalance"`
}

type seedRestaurantMenu struct {
	Name  string `json:"menu_name"`
	Price string `json:"price"`
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
			"INSERT INTO users (user_name, cash_balance) VALUES ($1, $2)",
			u.Name, u.CashBalance)
		if err != nil {
			return err
		}
	}
	return nil
}

func seedRestaurants(tx *sql.Tx) ([]int64, error) {
	b, err := os.ReadFile("seed/restaurant.json")
	if err != nil {
		return nil, err
	}

	var restaurant []seedRestaurant
	if err := json.Unmarshal(b, &restaurant); err != nil {
		return nil, err
	}

	var restaurantId []int64

	ctx := context.Background()
	for _, u := range restaurant {
		var restaurantResult modals.Restaurant
		row := tx.QueryRowContext(ctx,
			"INSERT INTO restaurant (restaurant_name, cash_balance) VALUES ($1, $2)",
			u.Name, u.CashBalance)
		if err := row.Scan(&restaurantResult.ID, &restaurantResult.Name, &restaurantResult.CashBalance); err != nil {
			return nil, err
		}
		restaurantId = append(restaurantId, restaurantResult.ID)
	}
	return restaurantId, nil
}

func seedRestaurantMenus(tx *sql.Tx) error {
	b, err := os.ReadFile("seed/restaurant_menu.json")
	if err != nil {
		return err
	}

	var restaurantMenu []seedRestaurantMenu
	if err := json.Unmarshal(b, &restaurantMenu); err != nil {
		return err
	}

	// var restaurantId []int64

	ctx := context.Background()
	for _, u := range restaurantMenu {
		var restaurantResult modals.Restaurant
		row := tx.QueryRowContext(ctx,
			"INSERT INTO restaurant_menu (menu_name, price,restaurant_id) VALUES ($1, $2, $3)",
			u.Name, u.Price)
		if err := row.Scan(&restaurantResult.ID, &restaurantResult.Name, &restaurantResult.CashBalance); err != nil {
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
	rId, err := seedRestaurants(tx)
	log.Println(len(rId))
	if err != nil {
		log.Fatal("Error seeding users: ", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	log.Println("✔ Seeding done")
}
