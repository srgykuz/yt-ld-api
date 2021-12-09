package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Amaimersion/yt-ld-api/config"
)

var (
	// ErrNoRow is returned when requested row doesn't exists.
	ErrNoRow = errors.New("no row")
)

// Open opens a database and connects to it.
// You should call this function only once.
func Open(env config.EnvConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable&connect_timeout=10",
		env.DBUser,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
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
