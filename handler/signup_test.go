package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestSignUpUser(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	result, err := signUpUser(database, "secret")

	if err != nil {
		t.Fatal(err)
	}

	if len(result.AccessToken) == 0 {
		t.Error("access token is empty")
	}
}
