package server

import (
	"google.golang.org/grpc"
)

func NewApiServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	api.RegisterApiServer(grpcServer, &apiServer{})
	return grpcServer
}
