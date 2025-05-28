package db

import (
	"service/order/internal/application/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ProductCode string
	Quantity    int64
	UnitPrice   float64
	OrderID     uint
}

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
	CreatedAt  int64
}

type Adapter struct {
	db *gorm.DB
}

func NewAdaptar(postgresUrl string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(postgresUrl + "?sslmode=disable"))

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, err
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id int64) (*domain.Order, error) {
	var orderEntity Order
	res := a.db.Preload("OrderItems").First(&orderEntity, id)

	if res.Error != nil {
		return nil, res.Error
	}
	var orderItems []domain.OrderItem

	for _, item := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt,
	}

	return &order, nil
}

func (a Adapter) Create(order *domain.Order) error {
	var orderItems []OrderItem

	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		})
	}

	newOrder := Order{
		OrderItems: orderItems,
		CustomerID: order.CustomerID,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
	}

	res := a.db.Create(&newOrder)

	if res.Error != nil {
		return res.Error
	}
	order.ID = int64(newOrder.ID)

	return nil
}
