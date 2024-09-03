package repo

import (
	"context"
	"seyhadot/gogrpc/pb"
)

type EchoRepo struct {
	pb.UnimplementedEchoServer
}

func NewEchoRepo() *EchoRepo {
	return &EchoRepo{}
}

func (e *EchoRepo) EchoBack(c context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.Message}, nil
}
