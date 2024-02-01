package main

import (
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"hexagon/adapters/postgrestore"
	"hexagon/pkg/config"
	"log"
	"strconv"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	db, err := postgrestore.NewConnection(postgrestore.Options{
		DBName:   cfg.DB.Name,
		DBUser:   cfg.DB.User,
		Password: cfg.DB.Pass,
		Host:     cfg.DB.Host,
		Port:     strconv.Itoa(cfg.DB.Port),
		SSLMode:  false,
	})
	if err != nil {
		log.Fatalf("cannot connecting to db: %v\n", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	total, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("cannot execute migration: %v\n", err)
	}

	log.Printf("applied %d migrations\n", total)
}
