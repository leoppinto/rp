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

	jobConfigurationServer := pb.NewJobResourceConfigurationServiceClient(conn)

	response, err := jobConfigurationServer.GetByJobId(context.Background(), &pb.GetByJobIdRequest{Jobid: "JOB-0001-13-01-2022"})
	if err != nil {
		log.Fatalf("Error when calling CreateProduct: %s", err)
	}
	log.Printf("Response from server: %v", response)

}
