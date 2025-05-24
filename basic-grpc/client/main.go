package main

import (
	"context"
	"fmt"

	pb "go-grpc/proto/productinfo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("hello gRPC client")
	conn, err := grpc.NewClient(":4005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	fmt.Print("conn", conn)
	client := pb.NewProductInfoClient(conn)

	productId, err := client.AddProduct(context.Background(), &pb.Product{
		Id:   "1",
		Name: "vishal1",
	})

	if err != nil {
		panic(err)
	}

	fmt.Print("productId", productId)

	product, err := client.GetProduct(context.Background(), productId)

	if err != nil {
		panic(err)
	}

	fmt.Print("product", product)
}
