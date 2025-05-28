package main

import (
	"fmt"
	"service/order/internal/adapters/db"
	"service/order/internal/application/core/domain"
)

func main() {
	adapter, err := db.NewAdaptar("postgres://postgres:mypassword@localhost:5433/postgres")

	if err != nil {
		panic(err)
	}

	order := domain.Order{
		CustomerID: 100,
		Status:     "Pending",
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "ProductCode1",
				Quantity:    100,
				UnitPrice:   100,
			},
			{
				ProductCode: "ProductCode2",
				Quantity:    200,
				UnitPrice:   200,
			},
		},
	}

	err = adapter.Create(&order)
	if err != nil {
		panic(err)
	}

	orderGet, err := adapter.Get(order.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("order", orderGet)
}
