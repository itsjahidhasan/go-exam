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
	Name        string  `json:"name"`
	CashBalance float64 `json:"cashBalance"`
}
type seedRestaurant struct {
	Name        string  `json:"name"`
	CashBalance float64 `json:"cashBalance"`
}

type seedRestaurantMenu struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func seedUsers(tx *sql.Tx) error {
	b, err := os.ReadFile("seed/user.json")
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

	var restaurantIds []int64
	ctx := context.Background()

	for _, r := range restaurant {
		var id int64
		// RETURNING id allows us to capture the generated ID
		err := tx.QueryRowContext(ctx,
			"INSERT INTO restaurant (restaurant_name, cash_balance) VALUES ($1, $2) RETURNING id",
			r.Name, r.CashBalance).Scan(&id)
		if err != nil {
			return nil, err
		}
		restaurantIds = append(restaurantIds, id)
	}
	return restaurantIds, nil
}

func seedRestaurantMenus(tx *sql.Tx, restaurantIds []int64) error {
	b, err := os.ReadFile("seed/restaurant_menu.json")
	if err != nil {
		return err
	}

	var restaurantMenu []seedRestaurantMenu
	if err := json.Unmarshal(b, &restaurantMenu); err != nil {
		return err
	}

	ctx := context.Background()
	for i, m := range restaurantMenu {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO restaurant_menu (menu_name, price, restaurant_id) VALUES ($1, $2, $3)",
			m.Name, m.Price, restaurantIds[i])
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
	rId, err := seedRestaurants(tx)
	log.Println(len(rId))
	if err != nil {
		log.Fatal("Error seeding Restaurants: ", err)
	}

	if err := seedRestaurantMenus(tx, rId); err != nil {
		log.Fatal("Error seeding Restaurant menus: ", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	log.Println("✔ Seeding done")
}
