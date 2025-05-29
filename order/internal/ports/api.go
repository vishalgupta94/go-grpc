package ports

import (
	"service/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(*domain.Order) error
}
