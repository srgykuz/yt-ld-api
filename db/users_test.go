package db_test

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestCreateUser(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	_, err = db.CreateUser(database)

	if err != nil {
		t.Error(err)
	}
}
