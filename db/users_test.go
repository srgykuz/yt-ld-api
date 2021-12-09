package db_test

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
)

func TestCreateUser(t *testing.T) {
	database, closeDB, err := openTestDB()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	_, err = db.CreateUser(database)

	if err != nil {
		t.Error(err)
	}
}
