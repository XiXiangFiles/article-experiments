package grpcserver

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Port       string
	GRPCServer *grpc.Server
	log        *zap.SugaredLogger
}

func NewGRPCServer(port string, log *zap.SugaredLogger) *GRPCServer {
	return &GRPCServer{
		Port:       port,
		GRPCServer: grpc.NewServer(),
		log:        log,
	}
}

func (s *GRPCServer) Run() {
	reflection.Register(s.GRPCServer)
	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		s.log.Warn(err)
	}
	s.log.Info("start gRPC server")
	if err := s.GRPCServer.Serve(lis); err != nil {
		s.log.Warn(err)
	}
}
