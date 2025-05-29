package gprc

import (
	"context"
	"service/order/internal/application/core/domain"
	pb "service/proto/order"
)

func (a Adapter) Create(ctx context.Context, in *pb.Order) (*pb.OrderResponse, error) {

	userId := in.UserId
	var orderItems []domain.OrderItem
	for _, orderItem := range in.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			Quantity:    orderItem.Quantity,
			UnitPrice:   float64(orderItem.UnitPrice),
		})
	}

	order := domain.NewOrder(userId, orderItems)
	err := a.api.PlaceOrder(&order)

	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		ResponseId: order.ID,
	}, nil
}
