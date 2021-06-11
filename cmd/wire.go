//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kouame-florent/axone-api/api/grpc/server"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/store"
	"github.com/kouame-florent/axone-api/internal/svc"
)

func InitializeNewAxoneServer() *server.AxoneServer {
	wire.Build(server.NewAxoneServer, svc.NewTicketSvc, repo.NewTicketRepo, store.NewDB)
	return &server.AxoneServer{}
}
