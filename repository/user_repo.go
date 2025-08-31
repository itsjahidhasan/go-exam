package repository

import (
	"context"
	"database/sql"
	"go-exam/modals"
	"log"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, u modals.User) error {
	q := "INSERT INTO users (user_name,cash_balance) VALUES ($1,$2) RETURNING id,cash_balance"

	return r.DB.QueryRowContext(ctx, q, u.Name, u.CashBalance).Scan(&u.ID)
}

func (r *UserRepo) GetAll(ctx context.Context) ([]modals.User, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id,name,cash_balance FROM users ORDER BY id DESC")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []modals.User
	for rows.Next() {
		var u modals.User
		if err := rows.Scan(&u.ID, &u.Name, &u.CashBalance); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*modals.User, error) {
	var u modals.User
	q := "SELECT id,name,cash_balance FROM users WHERE id=$1"
	err := r.DB.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.Name, &u.CashBalance)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, err
}

func (r *UserRepo) GetUserPurchaseHistory(ctx context.Context, id int64) ([]modals.UserPurchaseHistory, error) {
	q := `
	SELECT p.id,u.user_name,u.cash_balance,p.dish_id,p.amount,r.restaurant_name FROM 
	purchase_history p JOIN users u ON p.user_id = u.id 
	JOIN restaurant r ON p.restaurant_id = r.id WHERE p.user_id=$1
	`
	rows, err := r.DB.QueryContext(ctx, q, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var out []modals.UserPurchaseHistory
	for rows.Next() {
		var u modals.UserPurchaseHistory
		if err := rows.Scan(&u.ID, &u.Name, &u.CashBalance, &u.DishId, &u.DishPrice, u.RestaurantName); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}

func (r *UserRepo) PurchaseDish(ctx context.Context, d modals.UserPurchaseRequest) (modals.UserPurchaseResponse, error) {
	updateRestaurantBalanceQuery := `
		UPDATE restaurant SET cash_balance=cash_balance+$1 WHERE id=$2
	`
	updateUserBalanceQuery := `
		UPDATE users SET cash_balance=cash_balance - $1 WHERE id=$2
	`
	q := `
		INSERT INTO purchase_history (dish_id,restaurant_id,user_id,amount) 
		VALUES ($1,$2,$3,$4) 
		RETURNING id,dish_id,restaurant_id,user_id,amount
	`
	tx, err := r.DB.Begin()
	if err != nil {
		log.Println("Failed to start transaction")
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, updateRestaurantBalanceQuery)
	if err != nil {
		log.Println("Failed to exec query for updating restaurant cash_balance")
	}
	_, err = tx.ExecContext(ctx, updateUserBalanceQuery)
	if err != nil {
		log.Println("Failed to exec query for updating user cash_balance")
	}
	_, err = tx.ExecContext(ctx, q)
	if err != nil {
		log.Println("Failed to exec query for insert purchase_history")
	}

	if err := tx.Commit(); err != nil {
		log.Fatalln("Transaction Failed:", err)
	}

	return modals.UserPurchaseResponse{}, nil
}
