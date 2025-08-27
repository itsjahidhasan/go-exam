package modals

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Name      int64     `json:"name" db:"name"`
	Email     int64     `json:"email" db:"email"`
	Age       int       `json:"age" db:"age"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
