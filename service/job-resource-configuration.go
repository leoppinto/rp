package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"rp/grpc/pb"
	"rp/model"
	"time"
)

type JobResourceConfigurationGrpcServer struct {
	pb.UnimplementedJobResourceConfigurationServiceServer
	JobResouceConfigurations *model.JobResouceConfigurations
}

func (p *JobResourceConfigurationGrpcServer) GetByJobId(ctx context.Context, in *pb.GetByJobIdRequest) (*pb.GetByJobIdResponse, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Get_job_id_server_init_response")
	time.Sleep(time.Duration(2) * time.Millisecond)
	span.Finish()

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
