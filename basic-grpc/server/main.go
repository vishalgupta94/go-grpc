package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	pb "go-grpc/proto/productinfo"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV6()

	if err != nil {
		panic(err)
	}

	id := out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[id] = in
	return &pb.ProductID{Id: id}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, ok := s.productMap[in.Id]

	if !ok {
		return nil, errors.New("product not found")
	}
	return value, nil

}

func main() {
	fmt.Println("hello gRPC server")
	listener, err := net.Listen("tcp", ":4005")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductInfoServer(grpcServer, &server{})

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
