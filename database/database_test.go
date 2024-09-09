package database

import (
	"log"
	"testing"

	"github.com/codegram01/wingram-min/config"
)

func DbTest() *Db {
	cfg := config.Init()

	db, err := Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestConnect(t *testing.T) {
	cfg := config.Init()

	db, err := Connect(cfg)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(db)
}
