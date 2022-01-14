package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rp/grpc/pb"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("192.168.0.15:8300", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
		fmt.Println("Leonardo")
	}
	defer conn.Close()

	productServer := pb.NewProductServiceClient(conn)

	response, err := productServer.CreateProduct(context.Background(), &pb.Product{Name: "teste"})
	if err != nil {
		log.Fatalf("Error when calling CreateProduct: %s", err)
	}
	log.Printf("Response from server: %v", response)

}
