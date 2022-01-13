package rp

import (
	"log"
	"net"
	"rp/grpc/pb"
	"rp/model"
	"rp/service"

	"google.golang.org/grpc"
)

var ProductsList = model.NewProducts()

func StartServer() {

	lis, err := net.Listen("tcp", "192.168.0.15:8300")
	if err != nil {
		log.Print("Erro ao se conectar")
	}

	grpcServer := grpc.NewServer()
	productService := service.NewProductGrpcServer(ProductsList)
	pb.RegisterProductServiceServer(grpcServer, productService)

	log.Print("Servidor Conectado")

	_ = grpcServer.Serve(lis)
}
