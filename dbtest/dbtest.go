package dbtest

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Amaimersion/yt-ld-api/config"
	"github.com/Amaimersion/yt-ld-api/db"
)

var mu sync.Mutex

// Open opens DB that is intended for "one test" purpose.
//
// For DB credentials it reads .env.testing file at the root.
// Note that at moment of opening the DB already should be created
// (i.e. CREATE DATABASE).
//
// Opened DB will be migrated to actual state.
//
// Cleanup function will be returned. You must call this function
// when you done your test. If you will not call it, then next call
// to Open will block current thread. Even if Open returns error,
// you anyway should call cleanup function.
//
// Note that Open is safe for concurrent access only across single package.
// If multiple packages will call Open in parallel, then most likely you will
// get an error which tells that DB is up to date. You shouldn't call
// Open in parallel, limit to 1 thread only. For example, go test ./...
// by default runs multiple packages in parallel in case if you have enough
// CPU cores, in that case you should call with -p 1 flag.
//
// So, opened DB should be used only for one active test across all packages.
// Test life cycle: open -> populate -> test -> clean.
func Open() (*sql.DB, func() error, error) {
	mu.Lock()

	emptyCleanup := func() error {
		defer mu.Unlock()
		return nil
	}

	if err := config.LoadEnv("../.env.testing"); err != nil {
		return nil, emptyCleanup, err
	}

	env := config.ReadEnv()
	args := db.OpenArgs{
		User:     env.DBUser,
		Password: env.DBPassword,
		Host:     env.DBHost,
		Port:     env.DBPort,
		Name:     env.DBName,
	}
	database, err := db.Open(args)

	if err != nil {
		return nil, emptyCleanup, err
	}

	instance, err := postgres.WithInstance(
		database,
		&postgres.Config{},
	)

	if err != nil {
		return nil, emptyCleanup, err
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres",
		instance,
	)

	if err != nil {
		return nil, emptyCleanup, err
	}

	if err := migrations.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return nil, emptyCleanup, errors.New("DB is up to date. Did you not run cleanup at last open?")
		}

		return nil, emptyCleanup, err
	}

	cleanup := func() error {
		defer mu.Unlock()

		if err := migrations.Drop(); err != nil {
			return err
		}

		sourceErr, databaseErr := migrations.Close()

		if sourceErr != nil {
			return sourceErr
		}

		if databaseErr != nil {
			return databaseErr
		}

		return nil
	}

	return database, cleanup, nil
}
