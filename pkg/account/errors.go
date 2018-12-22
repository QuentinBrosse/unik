package account

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
	UnknownError       = status.Errorf(codes.Unknown, "something wrong happen, please try later")
	UserAlreadyExists  = status.Errorf(codes.AlreadyExists, "this user already exists")
)
