package ports

import "service/order/internal/application/core/domain"

// Ensure the 'domain' package exists and contains the 'Order' struct.

type DBPort interface {
	Get(id int64) (*domain.Order, error)
	Create(*domain.Order) error
}
