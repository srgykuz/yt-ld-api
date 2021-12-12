package handler

import (
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/dbtest"
)

func TestRemoveLike(t *testing.T) {
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

	if err := removeLike(database, args, userID); err != nil {
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

func TestRemoveLikeFirstTime(t *testing.T) {
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

	if err := removeLike(database, args, userID); err != nil {
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

func TestRemoveLikeMultipleTimes(t *testing.T) {
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

	if err := removeLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	if err := removeLike(database, args, userID); err != nil {
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

func TestRemoveLikeKeepsDislike(t *testing.T) {
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

	if err := removeLike(database, args, userID); err != nil {
		t.Fatal(err)
	}

	stat, err := getStat(database, args, userID)

	if err != nil {
		t.Fatal(err)
	}

	wantStat := getStatResult{
		LikesCount:    0,
		DislikesCount: 1,
		HasLike:       false,
		HasDislike:    true,
	}

	if stat != wantStat {
		t.Errorf("stat = %v, want = %v", stat, wantStat)
	}
}
