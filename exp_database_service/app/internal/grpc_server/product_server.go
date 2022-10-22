package grpcserver

import (
	factory "app/internal/database/collection_factory"
	"app/pb"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductGrpcServer struct {
	Factory factory.CollectionFactory
	Log     *zap.SugaredLogger
	pb.UnsafeExpServiceServer
}

func (p *ProductGrpcServer) AddData(ctx context.Context, param *pb.AddDataParameter) (*pb.AddDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddData not implemented")
}
func (p *ProductGrpcServer) QueryData(ctx context.Context, param *pb.QueryDataParamerter) (*pb.QueryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryData not implemented")
}
