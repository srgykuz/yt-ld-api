package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestRemoveDislike(t *testing.T) {
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

	if err := removeDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount:    0,
		DislikesCount: 0,
		HasLike:       false,
		HasDislike:    false,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestRemoveDislikeFirstTime(t *testing.T) {
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

	if err := removeDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount:    0,
		DislikesCount: 0,
		HasLike:       false,
		HasDislike:    false,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestRemoveDislikeMultipleTimes(t *testing.T) {
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

	if err := removeDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	if err := removeDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount:    0,
		DislikesCount: 0,
		HasLike:       false,
		HasDislike:    false,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}

func TestRemoveDislikeKeepsLike(t *testing.T) {
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

	if err := removeDislike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount:    1,
		DislikesCount: 0,
		HasLike:       true,
		HasDislike:    false,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}
