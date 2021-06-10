package main

import (
	"log"

	"github.com/kouame-florent/axone-api/internal/config"
	"github.com/kouame-florent/axone-api/internal/store"
)

func main() {
	dsn := config.DataSourceName()
	db := store.OpenDB(dsn)
	err := store.CreateSchema(db)
	if err != nil {
		log.Fatal(err)
	}

}
