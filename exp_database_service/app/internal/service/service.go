package service

import (
	"context"
	"fmt"

	factory "app/internal/database/collection_factory"
	"app/internal/global"
	server "app/internal/grpc_server"
	"app/pb"

	"go.uber.org/zap"
)

type Service struct {
	log *zap.SugaredLogger
	ctx context.Context
	svc *server.GRPCServer
}

func New(log *zap.SugaredLogger, ctx context.Context) *Service {
	s := server.NewGRPCServer(global.ServeAddress.ToString(), log)
	fmt.Println(global.MySQL_DSN.ToString())
	fmt.Println(global.MySQL_DATA_DB.ToString())
	mysql, err := factory.NewMySQLHandler(ctx, global.MySQL_DSN.ToString(), global.MySQL_DATA_DB.ToString())
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterExpServiceServer(s.GRPCServer, &server.ProductGrpcServer{Factory: mysql, Log: log})
	return &Service{
		log: log,
		ctx: ctx,
		svc: s,
	}
}

func (s *Service) Run() {
	s.svc.Run()
}
