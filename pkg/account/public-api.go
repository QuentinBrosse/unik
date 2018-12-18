package account

import (
	"context"

	"github.com/quentinbrosse/scwgame/protobufs/account"
)

type apiServer struct {
}

func (s *apiServer) SignUp(ctx context.Context, request *account.SignUpRequest) (*account.SignUpResponse, error) {
	reply := &account.SignUpResponse{
		Message: "Not implemented " + request.Email,
	}
	return reply, nil
}

func (s *apiServer) SignIn(ctx context.Context, request *account.SignInRequest) (*account.SignInResponse, error) {
	reply := &account.SignInResponse{
		Message: "Not implemented " + request.Email,
	}
	return reply, nil
}
