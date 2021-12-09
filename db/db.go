package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	// ErrNoRow is returned when requested row doesn't exists.
	ErrNoRow = errors.New("no row")
)

type OpenArgs struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// Open opens a database and connects to it.
// You should call this function only once.
func Open(args OpenArgs) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable&connect_timeout=10",
		args.User,
		args.Password,
		args.Host,
		args.Port,
		args.Name,
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
