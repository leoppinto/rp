package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"log"
	"rp/grpc/pb"
	"time"
)

func main() {

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

	// continue main()

	var conn *grpc.ClientConn

	conn, err = grpc.Dial(":8300", grpc.WithInsecure(),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer)),
		)))

	if err != nil {
		log.Fatalf("did not connect: %s", err)
		fmt.Println("Erro ao conectar")
	}
	defer conn.Close()

	jobConfigurationServer := pb.NewJobResourceConfigurationServiceClient(conn)

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "Get_job_id_client_init_request")
	time.Sleep(time.Duration(2) * time.Millisecond)
	span.Finish()

	response, err := jobConfigurationServer.GetByJobId(ctx, &pb.GetByJobIdRequest{Jobid: "JOB-0001-13-01-2022"})
	if err != nil {
		log.Fatalf("Error when calling CreateProduct: %s", err)
	}
	log.Printf("Response from server: %v", response)

}
