package service

import (
	"context"
	"rp/grpc/pb"
	"rp/model"
)

type JobResourceConfigurationGrpcServer struct {
	pb.UnimplementedJobResourceConfigurationServiceServer
	JobResouceConfigurations *model.JobResouceConfigurations
}

func (p *JobResourceConfigurationGrpcServer) GetByJobId(ctx context.Context, in *pb.GetByJobIdRequest) (*pb.GetByJobIdResponse, error) {
	return &pb.GetByJobIdResponse{
		Idrequest: 123,
		ResourceCreated: &pb.ResourceCreated{
			Id:       10000,
			Provider: "aws",
		},
	}, nil
}

func NewJobResourceConfigurationGrpcServer(jobResouceConfigurations *model.JobResouceConfigurations) *JobResourceConfigurationGrpcServer {
	return &JobResourceConfigurationGrpcServer{
		JobResouceConfigurations: jobResouceConfigurations,
	}

}
