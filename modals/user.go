package modals

type User struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"user_name"`
	CashBalance float64 `json:"cashBalance" db:"cash_balance"`
}

type UserPurchaseHistory struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	CashBalance float64 `json:"cashBalance" db:"cash_balance"`
	DishId      string  `json:"dishId" db:"dish_id"`
	DishPrice   string  `json:"dishPrice" db:"amount"`
}
