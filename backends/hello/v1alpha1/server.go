package main

import (
	"context"
	"net"
	"os"

	"go.alis.build/alog"
	"google.golang.org/grpc"

	pb "github.com/jaebrownn/proto-example/protobuf/go/hello/v1alpha1"
)

// HelloServiceServer is an implementation hello.v1alpha1.HelloServiceServer.
type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Greeting: "Hello " + req.Name}, nil
}

func (s *HelloServiceServer) SayGoodbye(ctx context.Context, req *pb.SayGoodbyeRequest) (*pb.SayGoodbyeResponse, error) {
	return &pb.SayGoodbyeResponse{Goodbye: "Goodbye " + req.Name}, nil
}

func main() {
	ctx := context.Background()

	// Get the port to run the server on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	alog.Noticef(ctx, "starting server on port "+port)

	// Set up a listener for the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		alog.Fatalf(ctx, "net.Listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// Attach server interceptors to the gRPC server
			serverInterceptor,
		),
	)
	pb.RegisterHelloServiceServer(grpcServer, &HelloServiceServer{})

	// Serve the gRPC server
	if err = grpcServer.Serve(listener); err != nil {
		alog.Fatal(ctx, err.Error())
	}
}

// serverInterceptor is an example of an interceptor which could be used to 'inject'
// for example logs and/or tracing details to incoming server requests, conduct custom authentication
// and authorization flows, etc.
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		alog.Debugf(ctx, "%v", req)
		alog.Warnf(ctx, err.Error())
	}

	_ = info
	return h, err
}
