package modals

type User struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"user_name"`
	CashBalance float64 `json:"cashBalance" db:"cash_balance"`
}

type UserPurchaseHistory struct {
	ID             int64   `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	CashBalance    float64 `json:"cashBalance" db:"cash_balance"`
	DishId         int64   `json:"dishId" db:"dish_id"`
	DishPrice      string  `json:"dishPrice" db:"amount"`
	RestaurantName string  `json:"restaurantName" db:"restaurant_name"`
}

type UserPurchaseRequest struct {
	UserId       int64  `json:"userId" db:"user_id"`
	DishId       int64  `json:"dishId" db:"dish_id"`
	RestaurantId int64  `json:"restaurantId" db:"restaurant_id"`
	Amount       string `json:"amount" db:"amount"`
}
type UserPurchaseResponse struct {
	ID           int64  `json:"id" db:"id"`
	UserId       int64  `json:"userId" db:"user_id"`
	DishId       int64  `json:"dishId" db:"dish_id"`
	RestaurantId int64  `json:"restaurantId" db:"restaurant_id"`
	Amount       string `json:"amount" db:"amount"`
}
