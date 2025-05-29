package main

import (
	"log"
	"service/order/internal/adapters/db"
	grpcAdapter "service/order/internal/adapters/grpc"
	"service/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdaptar("postgres://postgres:mypassword@localhost:5433/postgres")

	if err != nil {
		panic(err)
	}
	adpater := api.NewApplication(dbAdapter)
	log.Println(adpater)
	grpcAdapter := grpcAdapter.NewAdapter(adpater, 3000)
	grpcAdapter.Run()
	// order := domain.Order{
	// 	CustomerID: 100,
	// 	Status:     "Pending",
	// 	OrderItems: []domain.OrderItem{
	// 		{
	// 			ProductCode: "ProductCode1",
	// 			Quantity:    100,
	// 			UnitPrice:   100,
	// 		},
	// 		{
	// 			ProductCode: "ProductCode2",
	// 			Quantity:    200,
	// 			UnitPrice:   200,
	// 		},
	// 	},
	// }

	// err = adapter.Create(&order)
	// if err != nil {
	// 	panic(err)
	// }

	// orderGet, err := adapter.Get(order.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("order", orderGet)

}

// func main() {
// 	adapter, err := db.NewAdaptar("postgres://postgres:mypassword@localhost:5433/postgres")

// 	if err != nil {
// 		panic(err)
// 	}

// 	order := domain.Order{
// 		CustomerID: 100,
// 		Status:     "Pending",
// 		OrderItems: []domain.OrderItem{
// 			{
// 				ProductCode: "ProductCode1",
// 				Quantity:    100,
// 				UnitPrice:   100,
// 			},
// 			{
// 				ProductCode: "ProductCode2",
// 				Quantity:    200,
// 				UnitPrice:   200,
// 			},
// 		},
// 	}

// 	err = adapter.Create(&order)
// 	if err != nil {
// 		panic(err)
// 	}

// 	orderGet, err := adapter.Get(order.ID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("order", orderGet)

// }
