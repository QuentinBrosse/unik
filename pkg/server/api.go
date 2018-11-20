package server

import (
	"context"

	"github.com/quentinbrosse/unik/protobufs"
)

type apiServer struct {
}

func (s *apiServer) SayHello(ctx context.Context, request *api.HelloRequest) (*api.HelloReply, error) {
	reply := &api.HelloReply{
		Message: "Hello " + request.Name,
	}
	return reply, nil
}
