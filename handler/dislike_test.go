package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestSetDislike(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "test"
	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.CreateUserReactions(database, userID, videoID); err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		DislikesCount: 1,
		HasDislike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestSetDislikeFirstTime(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "test"
	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		DislikesCount: 1,
		HasDislike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestSetDislikeMultipleTimes(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "test"
	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	if err := setDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		DislikesCount: 1,
		HasDislike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestSetDislikeRemoveLike(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "test"
	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.CreateUserReactions(database, userID, videoID); err != nil {
		t.Fatal(err)
	}

	userReactions := db.UserReactions{
		UserID:     userID,
		VideoID:    videoID,
		HasLike:    true,
		HasDislike: false,
	}

	if err := db.UpdateUserReactions(database, userReactions); err != nil {
		t.Fatal(err)
	}

	if err := db.IncrementLikesCount(database, videoID); err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		DislikesCount: 1,
		HasDislike:    true,
		LikesCount:    0,
		HasLike:       false,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}
