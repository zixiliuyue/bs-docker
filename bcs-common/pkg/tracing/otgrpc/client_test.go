

package grpc

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	pb "github.com/Tencent/bk-bcs/bcs-common/pkg/tracing/otgrpc/hello"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestOpenTracingClientInterceptor(t *testing.T) {
	closer, err := initTracing(t, "client-grpc")
	if err != nil {
		t.Fatal()
		return
	}

	if closer != nil {
		defer closer.Close()
	}

	runHelloClient()
}

func runHelloClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
