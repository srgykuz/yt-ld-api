package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestGetStatNoSuchVideo(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	args := videoInfoArgs{
		VideoID: "test",
	}
	result, err := getStat(database, args, 1)

	if err != nil {
		t.Fatalf("err = %v, want = nil", err)
	}

	wantResult := getStatResult{
		LikesCount:    0,
		DislikesCount: 0,
		HasLike:       false,
		HasDislike:    false,
	}

	if result != wantResult {
		t.Errorf("result = %v, want = %v", result, wantResult)
	}
}

func TestGetStatNoSuchUser(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "test"
	err = db.IncrementLikesCount(database, videoID)

	if err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}
	result, err := getStat(database, args, 1)

	if err != nil {
		t.Fatalf("err = %v, want = nil", err)
	}

	wantResult := getStatResult{
		LikesCount:    1,
		DislikesCount: 0,
		HasLike:       false,
		HasDislike:    false,
	}

	if result != wantResult {
		t.Errorf("result = %v, want = %v", result, wantResult)
	}
}

func TestGetStat(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "test"
	err = db.IncrementLikesCount(database, videoID)

	if err != nil {
		t.Fatal(err)
	}

	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	err = db.CreateUserReactions(database, userID, videoID)

	if err != nil {
		t.Fatal(err)
	}

	userReactions := db.UserReactions{
		UserID:     userID,
		VideoID:    videoID,
		HasLike:    true,
		HasDislike: false,
	}
	err = db.UpdateUserReactions(database, userReactions)

	if err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}
	result, err := getStat(database, args, 1)

	if err != nil {
		t.Fatalf("err = %v, want = nil", err)
	}

	wantResult := getStatResult{
		LikesCount:    1,
		DislikesCount: 0,
		HasLike:       true,
		HasDislike:    false,
	}

	if result != wantResult {
		t.Errorf("result = %v, want = %v", result, wantResult)
	}
}
