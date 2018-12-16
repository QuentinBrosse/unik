package api

import (
	"context"

	"github.com/quentinbrosse/scwgame/protobufs/api"
)

type apiServer struct {
}

func (s *apiServer) SayHello(ctx context.Context, request *api.HelloRequest) (*api.HelloReply, error) {
	reply := &api.HelloReply{
		Message: "Hello " + request.Name,
	}
	return reply, nil
}
