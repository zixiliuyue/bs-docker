

package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"testing"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/tracing"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/tracing/jaeger"
	pb "github.com/Tencent/bk-bcs/bcs-common/pkg/tracing/otgrpc/hello"

	"github.com/opentracing/opentracing-go"
	tracinglog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func initTracing(t *testing.T, serviceName string) (io.Closer, error) {
	tracer, err := tracing.NewInitTracing(
		serviceName,
		tracing.TracerSwitch("on"),
		tracing.TracerType(tracing.Jaeger),
		tracing.ReportLog(true),
		tracing.SampleType(jaeger.SamplerTypeConst),
		tracing.SampleParameter(1))
	if err != nil {
		t.Fatal(err)
	}

	return tracer.Init()
}

func TestOpenTracingServerInterceptor(t *testing.T) {
	closer, err := initTracing(t, "server-grpc")
	if err != nil {
		t.Fatal()
		return
	}

	if closer != nil {
		defer closer.Close()
	}

	runGRPCServer()
}

// server is used to implement hello.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	formatString(ctx, "evan")
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func runGRPCServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(OpenTracingServerInterceptor(opentracing.GlobalTracer())),
	)

	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func formatString(ctx context.Context, helloTo string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	helloStr := fmt.Sprintf("hello, %s", helloTo)
	span.LogFields(
		tracinglog.String("event", "string-format"),
		tracinglog.String("value", helloStr),
	)
}
