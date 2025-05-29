package main

import (
	"context"
	"fmt"

	pb "service/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("hello gRPC client")
	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	fmt.Print("conn", conn)
	client := pb.NewOrdersClient(conn)

	orderResponse, err := client.Create(context.Background(), &pb.Order{
		UserId: 1122,
		OrderItems: []*pb.OrderItem{
			{
				ProductCode: "code1",
				UnitPrice:   11.22,
				Quantity:    64,
			},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Print("orderResponse", orderResponse)
}
