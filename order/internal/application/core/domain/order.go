package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_item"`
	Quantity    int64   `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type Order struct {
	ID         int64       `json:"id"` // only this to worry about
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {

	return Order{
		CustomerID: customerId,
		OrderItems: orderItems,
		Status:     "Pending",
		CreatedAt:  time.Now().Unix(),
	}
}
