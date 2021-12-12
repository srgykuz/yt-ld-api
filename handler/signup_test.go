package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestSignUpUser(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	result, err := signUpUser(database, "secret")

	if err != nil {
		t.Fatal(err)
	}

	if len(result.AccessToken) == 0 {
		t.Error("access token is empty")
	}
}
