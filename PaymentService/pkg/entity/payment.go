package entity

type Payment struct {
	OrderId    string `json:"order_id"`
	UserId     string `json:"user_id"`
	Status     string `json:"status"`
	TotalPrice int64  `json:"total_price"`
}
