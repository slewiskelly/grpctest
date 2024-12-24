// Package grpctest provides utilities for gRPC testing.
package grpctest

import (
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// A Server is an gRPC server for use in tests.
type Server struct {
	ln  net.Listener
	srv *grpc.Server

	mu       sync.RWMutex
	shutdown bool
}

// New initializes a (test) gRPC server with the provided options.
//
// The server listens on a system chosen port on `localhost`; use `Conn` for a
// connection to the server.
//
// All services to be served by the server must be provided via `Register`.
func New(opts ...Option) *Server {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(fmt.Sprintf("grpctest: failed to listen: %v", err))
	}

	s := &Server{
		ln:  ln,
		srv: grpc.NewServer(),
	}

	for _, fn := range opts {
		fn(s)
	}

	go s.srv.Serve(s.ln)

	return s
}

// Conn returns an (idle) connection to the server.
func (s *Server) Conn() *grpc.ClientConn {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.shutdown {
		panic("Server shutdown")
	}

	conn, err := grpc.NewClient(s.ln.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("grpctest: failed to connect to server: %v", err))
	}

	return conn
}

//	Stop immediately stops the server, not accepting any new requests.
//
// All open connections are immediately closed and all active RPCs are
// cancelled.
func (s *Server) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.shutdown = true

	s.srv.Stop()
}
