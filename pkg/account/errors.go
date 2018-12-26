package account

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	UnknownError       = status.Errorf(codes.Unknown, "something wrong happen, please try later")
	UserAlreadyExists  = status.Errorf(codes.AlreadyExists, "this user already exists")
	InvalidCredentials = status.Errorf(codes.PermissionDenied, "invalid credentials")
)
