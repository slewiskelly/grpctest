package grpctest

import "google.golang.org/grpc"

// Option is an option to a (test) gRPC server.
type Option func(*Server)

// Register registers  a service and its implementation to the server.
func Register[T any](fn func(s grpc.ServiceRegistrar, srv T), impl any) Option {
	return func(s *Server) {
		fn(s.srv, impl.(T))
	}
}
