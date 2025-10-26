package v1

import (
	"debez/pkg/contract"
	"debez/pkg/logger"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpc *grpc.Server
}

func NewServer(logger logger.Logger) *Server {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)

	return &Server{grpc: srv}
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	return s.grpc.Serve(lis)
}

func (s *Server) Stop() {
	s.grpc.GracefulStop()
}

func (s *Server) RegisterServices(srv contract.UserServiceServer) {
	contract.RegisterUserServiceServer(s.grpc, srv)
}
