package grpc

import (
	"service/internal/app/grpc/server"

	xtremegrpc "github.com/globalxtreme/go-core/v2/grpc"
)

func Register(srv *xtremegrpc.GRPCServer) {
	srv.Register(
		&server.TestingServer{},
		&server.SettingItemBrandServer{},
	)
}
