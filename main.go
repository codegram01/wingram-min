package main

import (
	"log"

	"github.com/codegram01/wingram-min/config"
	"github.com/codegram01/wingram-min/database"
	"github.com/codegram01/wingram-min/server"
)

func main() {
	cfg := config.Init()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server.Init(&server.ServerCfg{
		Cfg: cfg,
		Db:  db,
	})
}
