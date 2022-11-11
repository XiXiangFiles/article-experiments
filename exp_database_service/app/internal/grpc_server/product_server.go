package grpcserver

import (
	factory "app/internal/database/collection_factory"
	"app/internal/database/entities"
	"app/pb"
	"encoding/json"

	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductGrpcServer struct {
	CollectionName string
	Factory        factory.CollectionFactory
	Log            *zap.SugaredLogger
	pb.UnsafeExpServiceServer
}

func (p *ProductGrpcServer) AddData(ctx context.Context, param *pb.AddDataParameter) (*pb.AddDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddData not implemented")
}

func (p *ProductGrpcServer) QueryData(ctx context.Context, param *pb.QueryDataParameter) (*pb.QueryDataResponse, error) {
	db := p.Factory.GetCollection(p.CollectionName)
	dest := []*entities.ProductRaw{}
	id := int(param.Id)

	err := db.Query(ctx, &entities.ProductFilter{
		Id: &id,
	}, &dest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data, err := json.Marshal(dest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.QueryDataResponse{
		Id:   int32(id),
		Data: data,
	}, nil
}
