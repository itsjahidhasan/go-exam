package modals

type Restaurant struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"restaurant_name"`
	CashBalance float64 `json:"cashBalance" db:"cash_balance"`
}
