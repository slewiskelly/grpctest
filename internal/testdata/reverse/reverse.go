package reverse

import (
	"context"

	pb "github.com/slewiskelly/grpctest/internal/testdata/api/reverse/v1"
)

type Server struct {
	pb.UnimplementedReverseServiceServer
}

func (s *Server) Reverse(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: rev(req.GetMessage()),
	}, nil
}

func rev(s string) string {
	b := make([]byte, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		b[len(s)-i-1] = s[i]
	}

	return string(b)
}
