package main

import (
	"log"
	"net"
	"os"

	"github.com/quentinbrosse/scwgame/pkg/auth"
)

const serviceHost = "0.0.0.0"
const servicePort = "8010"

func main() {
	log.Printf("starting server on port %v", servicePort)

	serviceAddress := net.JoinHostPort(serviceHost, servicePort)
	lis, lisErr := net.Listen("tcp", serviceAddress)
	if lisErr != nil {
		log.Fatalf("failed to listen: %v", lisErr)
	}

	config := &auth.ServerConfig{
		SslCert: os.Getenv("SSL_CERT"),
		SslKey:  os.Getenv("SSL_KEY"),
	}

	server := auth.NewPublicApiServer(config)
	serverErr := server.Serve(lis)
	if serverErr != nil {
		log.Fatal(serverErr)
	}

	// TODO: GracefulStop
}
