package db_test

import (
	"database/sql"
	"testing"

	"github.com/Amaimersion/yt-ld-api/db"
)

func TestReadReaction(t *testing.T) {
	database, closeDB, err := openTestDB()

	if err != nil {
		t.Fatal(err)
	}

	defer closeDB()

	videoID := "random-video-id"

	// TODO: if db.CreateReaction will be implemented someday,
	// then use it first and then compare readed data with expected data
	_, err = db.ReadReaction(database, videoID)

	if err != db.ErrNoRow {
		t.Error(err)
	}
}

func TestIncrementDecrement(t *testing.T) {
	type test struct {
		name               string
		method             func(*sql.DB, string) error
		wantFirstReaction  db.Reaction
		wantSecondReaction db.Reaction
	}

	videoID := "random-video-id"
	tests := []test{
		{
			name:   "IncrementLikesCount",
			method: db.IncrementLikesCount,
			wantFirstReaction: db.Reaction{
				VideoID:    videoID,
				LikesCount: 1,
			},
			wantSecondReaction: db.Reaction{
				VideoID:    videoID,
				LikesCount: 2,
			},
		},
		{
			name:   "IncrementDislikesCount",
			method: db.IncrementDislikesCount,
			wantFirstReaction: db.Reaction{
				VideoID:       videoID,
				DislikesCount: 1,
			},
			wantSecondReaction: db.Reaction{
				VideoID:       videoID,
				DislikesCount: 2,
			},
		},
		{
			name:   "DecrementLikesCount",
			method: db.DecrementLikesCount,
			wantFirstReaction: db.Reaction{
				VideoID:    videoID,
				LikesCount: 0,
			},
			wantSecondReaction: db.Reaction{
				VideoID:    videoID,
				LikesCount: -1,
			},
		},
		{
			name:   "DecrementDislikesCount",
			method: db.DecrementDislikesCount,
			wantFirstReaction: db.Reaction{
				VideoID:       videoID,
				DislikesCount: 0,
			},
			wantSecondReaction: db.Reaction{
				VideoID:       videoID,
				DislikesCount: -1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(runT *testing.T) {
			database, closeDB, err := openTestDB()

			if err != nil {
				runT.Fatal(err)
			}

			defer closeDB()

			err = test.method(database, videoID)

			if err != nil {
				runT.Fatal(err)
			}

			firstReaction, err := db.ReadReaction(database, videoID)

			if err != nil {
				runT.Fatal(err)
			}

			if firstReaction != test.wantFirstReaction {
				runT.Fatalf("first reaction = %v, want = %v", firstReaction, test.wantFirstReaction)
			}

			err = test.method(database, videoID)

			if err != nil {
				runT.Fatal(err)
			}

			secondReaction, err := db.ReadReaction(database, videoID)

			if err != nil {
				runT.Fatal(err)
			}

			if secondReaction != test.wantSecondReaction {
				runT.Fatalf("second reaction = %v, want = %v", secondReaction, test.wantSecondReaction)
			}
		})
	}
}
