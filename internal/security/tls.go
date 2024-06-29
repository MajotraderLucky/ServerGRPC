package security

import (
	"github.com/MajotraderLucky/ServerGRPC/internal/config"
	"google.golang.org/grpc/credentials"
)

func SetupTLS(cfg *config.Config) (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile(cfg.ServerCert, cfg.ServerKey)
}
