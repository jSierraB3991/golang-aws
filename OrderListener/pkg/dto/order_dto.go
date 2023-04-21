package dto

type Order struct {
	OrderID    string `json:"order_id"`
	UserID     string `json:"user_id"`
	TotalPrice int64  `json:"total_price"`
	Status     string `json:"status"`
}
