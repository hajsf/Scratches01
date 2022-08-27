package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "streem/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {

	gRPcPort := ":50005"
	// Create a gRPC listener on TCP port
	lis, err := net.Listen("tcp", gRPcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterStreamServiceServer(s, server{})

	log.Println("Serving gRPC server on 0.0.0.0:50005")
	grpcTerminated := make(chan struct{})
	// Serve gRPC server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			close(grpcTerminated) // In case server is terminated without us requesting this
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	//	gateWayTarget := fmt.Sprintf("0.0.0.0%s", gRPcPort)
	conn, err := grpc.DialContext(
		context.Background(),
		gRPcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()

	log.Println("Serving gRPC-Gateway on http://localhost:8090")

	fmt.Println("run POST request of: http://localhost:8090/v1/example/echo with JSON data {'name': ' hello'}")
	fmt.Println("or run curl -X POST -k http://localhost:8090/v1/example/echo -d \"{'name': ' hello'}\"")

	log.Fatalln(gwServer.ListenAndServe())
}
