package service

import (
	"context"

	"github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"github.com/MajotraderLucky/ServerGRPC/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type simpleServer struct {
	pb.UnimplementedSimpleServiceServer
}

func (s *simpleServer) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: "Echo: " + in.GetMessage()}, nil
}

// CreateGRPCServer инициализирует gRPC сервер с TLS и JWT аутентификацией
func CreateGRPCServer(creds credentials.TransportCredentials) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.UnaryInterceptor(security.UnaryInterceptor()), // Добавляем JWT интерсептор
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSimpleServiceServer(grpcServer, &simpleServer{}) // Используем изменённое имя структуры
	return grpcServer
}
