package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"hexagon/adapters/httpserver"
	"hexagon/adapters/postgrestore"
	"hexagon/pkg/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgrestore.NewConnection(postgrestore.ParseFromConfig(cfg))
	if err != nil {
		log.Fatal(err)
	}

	server, err := httpserver.New()
	if err != nil {
		log.Fatal(err)
	}

	//db, err := inmemstore.NewConnection()
	server.BookStore = postgrestore.NewBookStore(db)
	//server.BookStore = inmemstore.NewBookStore(db)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Println("server started!")
	log.Fatal(http.ListenAndServe(addr, server))
}
