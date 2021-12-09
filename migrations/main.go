package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Amaimersion/yt-ld-api/config"
)

func main() {
	flagCfg := config.ReadFlags()

	if len(flagCfg.EnvFile) > 0 {
		if err := config.LoadEnv(flagCfg.EnvFile); err != nil {
			log.Fatalln(err)
		}
	}

	envCfg := config.ReadEnv()
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envCfg.DBUser,
		envCfg.DBPassword,
		envCfg.DBHost,
		envCfg.DBPort,
		envCfg.DBName,
	)
	m, err := migrate.New("file://migrations", dbURL)

	if err != nil {
		log.Fatalln(err)
	}

	if err := m.Up(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Done")
}
