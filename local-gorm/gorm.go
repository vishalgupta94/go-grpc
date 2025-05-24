package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ProductCode string
	Quantity    int64
	UnitPrice   int64
	OrderID     uint
}

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:mypassword@localhost:5433/postgres?sslmode=disable"))

	if err != nil {
		panic(err)
	}

	fmt.Println("hello world", db)
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		panic(err)
	}

	res := db.Create(&Order{
		CustomerID: 100,
		Status:     "Pending",
		OrderItems: []OrderItem{
			OrderItem{
				ProductCode: "ProductCode",
				Quantity:    200,
				UnitPrice:   200,
			},
		},
	})

	fmt.Println("res", res)

}
