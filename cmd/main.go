package main

import (
	"log"
	"os"

	"github.com/kouame-florent/axone-api/api/grpc/server"
	"github.com/kouame-florent/axone-api/internal/config"
	"github.com/kouame-florent/axone-api/internal/store"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	zap.ReplaceGlobals(logger)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	config.InitEnv(home)
	err = config.CreateAttachmentFolder(home)
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer store.CloseDB(db)

	store.CreateSchema(db) //only for test must be removed in production

	//buildSqlSchema(db)

	opts, err := server.ServerOptions()
	if err != nil {
		log.Fatal(err)
	}
	server.StartGrpcServer(opts, db)

}
