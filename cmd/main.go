package main

import (
	"log"
	"net"

	"github.com/kouame-florent/axone-api/api/grpc/gen"
	"github.com/kouame-florent/axone-api/api/grpc/server"
	"github.com/kouame-florent/axone-api/internal/config"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/store"
	"github.com/kouame-florent/axone-api/internal/svc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	buildSqlSchema()
	StartGrpcServer()

}

func buildSqlSchema() {
	dsn := config.DataSourceName()
	db := store.OpenDB(dsn)
	err := store.CreateSchema(db)
	if err != nil {
		log.Fatal(err)
	}
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := store.NewDB()
	tkRep := repo.NewTicketRepo(db)
	tkScv := svc.NewTicketSvc(tkRep)
	axSvr := server.NewAxoneServer(tkScv)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	gen.RegisterAxoneServer(s, axSvr)

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
