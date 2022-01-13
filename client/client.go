package main

import (
	"context"
	"log"
	"rp/grpc/pb"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8300", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	productServer := pb.NewProductServiceClient(conn)

	response, err := productServer.CreateProduct(context.Background(), &pb.Product{Name: "teste"})
	if err != nil {
		log.Fatalf("Error when calling CreateProduct: %s", err)
	}
	log.Printf("Response from server: %v", response)

}
