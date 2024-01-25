package domain

import "time"

type Orders struct {
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

type OrderItems struct {
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
