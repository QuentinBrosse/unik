package main

import (
	"log"
	"net"
	"os"

	"github.com/quentinbrosse/unik/pkg/server"
)

const defaultServiceAddress = "0.0.0.0:8020"

func main() {

	// print env
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	lis, lisErr := net.Listen("tcp", defaultServiceAddress)
	if lisErr != nil {
		log.Fatalf("failed to listen: %v", lisErr)
	}

	server := server.NewApiServer()
	serverErr := server.Serve(lis)
	if serverErr != nil {
		log.Fatal(serverErr)
	}

	// TODO: GracefulStop
}
