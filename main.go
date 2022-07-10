package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/italorfeitosa/go-grpc-lab/proto"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterQuotesServer(grpcServer, NewHandler())
	log.Println("grpc server running in ", fmt.Sprintf("localhost:%d", *port))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
