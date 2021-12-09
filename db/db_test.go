package db_test

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Amaimersion/yt-ld-api/config"
	"github.com/Amaimersion/yt-ld-api/db"
)

// openTestDB constructs DB that is intended for "one test" purpose.
//
// For DB credentials it reads .env.testing file at the root.
// Note that at moment of opening the DB already should be created
// (i.e. CREATE DATABASE).
//
// Opened DB will be migrated to actual state.
//
// Cleanup function will be returned. You must call this function
// when you done your test. In case if you will not call it, then
// at next openTestDB you will get an error, and in that case you
// will have to manually delete DB (i.e. DROP DATABASE). In case
// if openTestDB returns error, you can skip calling of cleanup.
//
// So, opened DB should be applied only for one test and no more!
// Life cycle: open -> populate (if needed) -> test -> clean
func openTestDB() (*sql.DB, func() error, error) {
	emptyCleanup := func() error {
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
