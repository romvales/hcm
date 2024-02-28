package grpcServer

import (
	"goServer/internal/core"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	GO_SERVER = "localhost:7261"
)

func init() {
	if val, exists := os.LookupEnv("GO_SERVER"); exists {
		GO_SERVER = val
	}
}

func StartGRPCServer() {
	var _netListenerError error
	var listener net.Listener
	var opts []grpc.ServerOption

	if listener, _netListenerError = net.Listen("tcp", GO_SERVER); _netListenerError != nil {
		log.Fatal(_netListenerError)
	}

	grpcServer := grpc.NewServer(opts...)

	defer listener.Close()
	defer grpcServer.Stop()

	core.Register(grpcServer)

	if _grpcServerErr := grpcServer.Serve(listener); _grpcServerErr != nil {
		log.Fatal(_grpcServerErr)
	}
}
