package api

import (
	"service/order/internal/application/core/domain"
	"service/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db,
	}
}

func (a *Application) PlaceOrder(order *domain.Order) error {
	err := a.db.Create(order)

	if err != nil {
		return err
	}
	return nil
}
