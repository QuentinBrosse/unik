package account

import (
	"context"

	pb "github.com/quentinbrosse/scwgame/protobufs/account"
)

type apiServer struct {
}

func (s *apiServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	account := &Account{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}

	err := account.Create()
	if err != nil {
		return nil, err
	}

	reply := &pb.SignUpResponse{
		Account: account.ToPb(),
	}
	return reply, nil
}

func (s *apiServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	reply := &pb.SignInResponse{}
	return reply, nil
}
