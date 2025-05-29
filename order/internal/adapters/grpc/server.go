package gprc

import (
	"log"
	"net"
	"service/order/internal/ports"
	"service/proto/order"

	"google.golang.org/grpc"
)

type Adapter struct {
	api  ports.APIPort
	port int64
	order.UnimplementedOrdersServer
}

func NewAdapter(api ports.APIPort, port int64) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		panic(err)
	}
	log.Printf("error  1 %v", err)
	grpcServer := grpc.NewServer()

	order.RegisterOrdersServer(grpcServer, a)
	err = grpcServer.Serve(listener)
	log.Printf("error 2 %v", err)
	if err != nil {
		panic(err)
	}
	log.Println("grpc server up and running")
}
