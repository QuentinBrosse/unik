package main

import (
	"log"
	"net"
	"os"

	"github.com/quentinbrosse/scwgame/pkg/api"
)

const serviceHost = "0.0.0.0"
const servicePort = "8020"

func main() {
	log.Printf("starting server on port %v", servicePort)

	serviceAddress := net.JoinHostPort(serviceHost, servicePort)
	lis, lisErr := net.Listen("tcp", serviceAddress)
	if lisErr != nil {
		log.Fatalf("failed to listen: %v", lisErr)
	}

	config := &api.ServerConfig{
		SslCert: os.Getenv("SSL_CERT"),
		SslKey:  os.Getenv("SSL_KEY"),
	}

	server := api.NewApiServer(config)
	info := server.GetServiceInfo()
	log.Println("GetServiceInfo:", len(info))
	for k, v := range info {
		log.Println(k, v)
	}
	serverErr := server.Serve(lis)
	if serverErr != nil {
		log.Fatal(serverErr)
	}

	// TODO: GracefulStop
}
