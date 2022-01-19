package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"log"
	"net"
	"rp/grpc/pb"
	"rp/model"
	"rp/service"
)

var ProductsList = model.NewProducts()
var JobConfigurationList = model.NewJobResouceConfigurations()

func StartServer() {

	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "resource_provisioner",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	lis, err := net.Listen("tcp", ":8300")
	if err != nil {
		log.Print("[ERROR] Grpc Server not connected")
	}

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)))

	productService := service.NewProductGrpcServer(ProductsList)
	pb.RegisterProductServiceServer(grpcServer, productService)
	jobConfigurationService := service.NewJobResourceConfigurationGrpcServer(JobConfigurationList)
	pb.RegisterJobResourceConfigurationServiceServer(grpcServer, jobConfigurationService)
	log.Print("[INFO] Server Gprc Connected")
	log.Print("[INFO] Job Configuration Service Started")
	log.Print("[INFO] Product Service Started")
	grpcServer.Serve(lis)
}
