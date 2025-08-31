package repository

import (
	"context"
	"database/sql"
	"go-exam/modals"
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
	rows, err := r.DB.QueryContext(ctx, "SELECT (p.id,u.name,u.cash_balance,p.dish_id,p.amount) FROM purchase_history p JOIN users u on p.user_id = u.id WHERE p.user_id=$1", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []modals.UserPurchaseHistory
	for rows.Next() {
		var u modals.UserPurchaseHistory
		if err := rows.Scan(&u.ID, &u.Name, &u.CashBalance, &u.DishId, &u.DishPrice); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}
