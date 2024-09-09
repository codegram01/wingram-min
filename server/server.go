package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegram01/wingram-min/config"
	"github.com/codegram01/wingram-min/database"
)

type ServerCfg struct {
	Cfg *config.Config
	Db  *database.Db
}

type Server struct {
	Db *database.Db
	Mux *http.ServeMux
}

func Init(scfg *ServerCfg) {
	mux := http.NewServeMux()

	s := &Server{
		Db: scfg.Db,
		Mux: mux,
	}

	s.makeHandler()

	addr := fmt.Sprintf(":%s", scfg.Cfg.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("server running ", addr)
	log.Fatal(server.ListenAndServe())
}
