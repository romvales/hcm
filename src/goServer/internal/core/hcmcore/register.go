package hcmcore

import (
	"goServer/internal/core/pb"

	"google.golang.org/grpc"
)

func Register(srv *grpc.Server) {
	pb.RegisterCoreServiceServer(srv, NewCoreServiceServer())
}
