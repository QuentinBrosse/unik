package server

import (
	"github.com/quentinbrosse/unik/protobufs"
	"google.golang.org/grpc"
)

func NewApiServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	api.RegisterApiServer(grpcServer, &apiServer{})
	return grpcServer
}
