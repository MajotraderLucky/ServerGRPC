package main

import (
	"context"
	"log"
	"net"

	"github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"github.com/MajotraderLucky/ServerGRPC/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSimpleServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: "Echo: " + in.GetMessage()}, nil
}

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Setup TLS
	creds, err := credentials.NewServerTLSFromFile(cfg.ServerCert, cfg.ServerKey)
	if err != nil {
		log.Fatalf("failed to load TLS credentials: %v", err)
	}

	// Create gRPC server with TLS
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSimpleServiceServer(s, &server{})

	lis, err := net.Listen("tcp", cfg.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
