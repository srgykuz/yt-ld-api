package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Amaimersion/yt-alt-ld-api/config"
)

// Open opens a database and connects to it.
// You should call this function only once.
func Open() (*sql.DB, error) {
	cfg := config.ReadEnv()
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable&connect_timeout=10",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
