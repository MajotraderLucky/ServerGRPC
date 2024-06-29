package service

import (
	"context"

	"github.com/MajotraderLucky/ServerGRPC/api/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSimpleServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: "Echo: " + in.GetMessage()}, nil
}

func CreateGRPCServer(creds credentials.TransportCredentials) *grpc.Server {
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSimpleServiceServer(s, &server{})
	return s
}
