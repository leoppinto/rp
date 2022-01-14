package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"rp/grpc/pb"
	"rp/model"
	"rp/service"
)

var ProductsList = model.NewProducts()
var JobConfigurationService = model.NewJobResouceConfigurations()

func StartServer() {
	lis, err := net.Listen("tcp", "192.168.0.15:8300")
	if err != nil {
		log.Print("[ERROR] Grpc Server not connected")
	}
	grpcServer := grpc.NewServer()
	productService := service.NewProductGrpcServer(ProductsList)
	pb.RegisterProductServiceServer(grpcServer, productService)
	jobConfigurationService := service.NewJobResourceConfigurationGrpcServer(JobConfigurationService)
	pb.RegisterJobResourceConfigurationServiceServer(grpcServer, jobConfigurationService)
	log.Print("[INFO] Server Gprc Connected")
	log.Print("[INFO] Job Configuration Service Started")
	log.Print("[INFO] Product Service Started")
	grpcServer.Serve(lis)
}
