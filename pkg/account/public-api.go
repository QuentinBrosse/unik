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

func (s *apiServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	account := &Account{
		Email:    req.Email,
		Password: req.Password,
	}

	err := account.SignIn()
	if err != nil {
		return nil, err
	}

	reply := &pb.SignInResponse{
		Account: account.ToPb(),
	}
	return reply, nil
}
