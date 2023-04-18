package entity

type Payment struct {
	OrderId    string `json:"order_id"`
	Status     string `json:"status"`
	TotalPrice int64  `json:"total_price"`
}
