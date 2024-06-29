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
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	creds, err := setupTLS(cfg)
	if err != nil {
		log.Fatalf("failed to load TLS credentials: %v", err)
	}

	gRPCServer := createGRPCServer(creds)
	startServer(gRPCServer, cfg.ServerAddress)
}

func loadConfig() (*config.Config, error) {
	return config.LoadConfig("config/config.json")
}

func setupTLS(cfg *config.Config) (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile(cfg.ServerCert, cfg.ServerKey)
}

func createGRPCServer(creds credentials.TransportCredentials) *grpc.Server {
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSimpleServiceServer(s, &server{})
	return s
}

func startServer(s *grpc.Server, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
