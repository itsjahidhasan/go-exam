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
	q := "INSERT INTO users (name,email,age) VALUES ($1,$2,$3) RETURNING id,created_at"

	return r.DB.QueryRowContext(ctx, q, u.Name, u.Email, u.Age).Scan(&u.ID, &u.CreatedAt)
}

func (r *UserRepo) GetAll(ctx context.Context) ([]modals.User, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id,name,email,age,created_at FROM users ORDER BY created_at DESC")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []modals.User
	for rows.Next() {
		var u modals.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*modals.User, error) {
	var u modals.User
	q := "SELECT id,name,email,age,created_at FROM users WHERE id=$1"
	err := r.DB.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, err
}

func (r *UserRepo) Update(ctx context.Context, u *modals.User) (*modals.User, error) {
	q := "UPDATE users SET name=$1,email=$2,age=$3 WHERE id=$4 RETURNING id,name,email,age,created_at"
	var out modals.User
	err := r.DB.QueryRowContext(ctx, q, u.Name, u.Email, u.Age, u.ID).Scan(&out.ID, &out.Name, &out.Email, &out.Age, &out.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) (bool, error) {
	res, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)

	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
