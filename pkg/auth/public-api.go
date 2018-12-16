package auth

import (
	"context"

	"github.com/quentinbrosse/scwgame/protobufs/auth"
)

type apiServer struct {
}

func (s *apiServer) SignUp(ctx context.Context, request *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	reply := &auth.SignUpResponse{
		Message: "Not implemented " + request.Email,
	}
	return reply, nil
}

func (s *apiServer) SignIn(ctx context.Context, request *auth.SignInRequest) (*auth.SignInResponse, error) {
	reply := &auth.SignInResponse{
		Message: "Not implemented " + request.Email,
	}
	return reply, nil
}
