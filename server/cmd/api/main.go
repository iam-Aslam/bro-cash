package main

import (
	"log"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/config"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatalf("failed initialize api: %v\n", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
