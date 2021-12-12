package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestSetLike(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

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

	if err := setLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount: 1,
		HasLike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestSetLikeFirstTime(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	videoID := "test"
	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount: 1,
		HasLike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestSetLikeMultipleTimes(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

	videoID := "test"
	userID, err := db.CreateUser(database)

	if err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	if err := setLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount: 1,
		HasLike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestSetLikeRemoveDislike(t *testing.T) {
	database, closeDB, err := dbtest.Open()

	defer closeDB()

	if err != nil {
		t.Fatal(err)
	}

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
		HasLike:    false,
		HasDislike: true,
	}

	if err := db.UpdateUserReactions(database, userReactions); err != nil {
		t.Fatal(err)
	}

	if err := db.IncrementDislikesCount(database, videoID); err != nil {
		t.Fatal(err)
	}

	args := videoInfoArgs{
		VideoID: videoID,
	}

	if err := setLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		DislikesCount: 0,
		HasDislike:    false,
		LikesCount:    1,
		HasLike:       true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}
