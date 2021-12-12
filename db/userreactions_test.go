package db_test

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestCreateUserReactions(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	videoID := "random-video-id"
	err = db.CreateUserReactions(database, userID, videoID)

	if err != nil {
		t.Error(err)
	}
}

func TestReadUserReactions(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	videoID := "random-video-id"
	err = db.CreateUserReactions(database, userID, videoID)

	if err != nil {
		t.Fatal(err)
	}

	_, err = db.ReadUserReactions(database, userID, videoID)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUserReactions(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	videoID := "random-video-id"
	err = db.CreateUserReactions(database, userID, videoID)

	if err != nil {
		t.Fatal(err)
	}

	wantUserReactions := db.UserReactions{
		UserID:     userID,
		VideoID:    videoID,
		HasLike:    true,
		HasDislike: false,
	}
	err = db.UpdateUserReactions(database, wantUserReactions)

	if err != nil {
		t.Fatal(err)
	}

	userReactions, err := db.ReadUserReactions(database, userID, videoID)

	if err != nil {
		t.Fatal(err)
	}

	if wantUserReactions != userReactions {
		t.Error("expected result not equal to actual result")
	}
}
