package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	m, err := migrate.New("file://migrations", dbURL)

	if err != nil {
		log.Fatalln(err)
	}

	if err := m.Up(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Done")
}
