package main

import (
	"context"
	"log"
	"net"

	"github.com/quii/go-specs-greet/adapters/grpcserver"
	"github.com/quii/go-specs-greet/domain/interactions"
	"google.golang.org/grpc"
)

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Greet(context.Context, *grpcserver.GreetRequest) (*grpcserver.GreetReply, error)
	mustEmbedUnimplementedGreeterServer()
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: interactions.Greet(request.Name)}, nil
}
