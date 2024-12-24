package echo

import (
	"context"

	pb "github.com/slewiskelly/grpctest/internal/testdata/api/echo/v1"
)

type Server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *Server) Echo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: req.GetMessage(),
	}, nil
}
