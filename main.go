package main

import (
	"os"

	"github.com/GolfZaaa/Golang-ProjectE-commerce/config"
	"github.com/GolfZaaa/Golang-ProjectE-commerce/modules/servers"
	"github.com/GolfZaaa/Golang-ProjectE-commerce/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
