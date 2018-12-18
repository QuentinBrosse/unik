package account

import (
	"github.com/quentinbrosse/scwgame/pkg/std/tls"
	"github.com/quentinbrosse/scwgame/protobufs/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServerConfig struct {
	SslCert string
	SslKey  string
}

func NewPublicApiServer(config *ServerConfig) *grpc.Server {
	cert := tls.LoadCertif(config.SslCert, config.SslKey)

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	server := grpc.NewServer(opts...)
	account.RegisterPublicApiServer(server, &apiServer{})
	return server
}
