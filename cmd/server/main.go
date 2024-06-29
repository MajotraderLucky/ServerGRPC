package main

import (
	"log"

	"github.com/MajotraderLucky/ServerGRPC/internal/config"
	"github.com/MajotraderLucky/ServerGRPC/internal/service"
	"google.golang.org/grpc/credentials"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	creds, err := setupTLS(cfg)
	if err != nil {
		log.Fatalf("failed to load TLS credentials: %v", err)
	}

	gRPCServer := service.CreateGRPCServer(creds)
	service.StartServer(gRPCServer, cfg.ServerAddress)
}

func loadConfig() (*config.Config, error) {
	return config.LoadConfig("config/config.json")
}

func setupTLS(cfg *config.Config) (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile(cfg.ServerCert, cfg.ServerKey)
}
