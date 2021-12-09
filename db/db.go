package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Amaimersion/yt-ld-api/config"
)

var (
	// ErrNoRow is returned when row that was requested doesn't exists.
	ErrNoRow = errors.New("no such row")
)

// Open opens a database and connects to it.
// You should call this function only once.
func Open(cfg config.EnvConfig) (*sql.DB, error) {
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
