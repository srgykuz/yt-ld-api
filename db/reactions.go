package db

import (
	"database/sql"
	"fmt"
)

// IncrementLikesCount increments likes count of specific video by 1.
// It creates a new row if such video_id doesn't exists.
func IncrementLikesCount(database *sql.DB, video_id string) error {
	_, err := database.Exec(
		"INSERT INTO reactions (video_id, likes_count, dislikes_count) VALUES ($1, 1, 0) ON CONFLICT (video_id) DO UPDATE SET likes_count = reactions.likes_count + 1",
		video_id,
	)

	if err != nil {
		return fmt.Errorf("IncrementLikesCount: %v", err)
	}

	return nil
}

// DecrementLikesCount decrements likes count of specific video by 1.
// It creates a new row if such video_id doesn't exists.
//
// Note that it doesn't checks if new value will be less than 0.
func DecrementLikesCount(database *sql.DB, video_id string) error {
	_, err := database.Exec(
		"INSERT INTO reactions (video_id, likes_count, dislikes_count) VALUES ($1, 0, 0) ON CONFLICT (video_id) DO UPDATE SET likes_count = reactions.likes_count - 1",
		video_id,
	)

	if err != nil {
		return fmt.Errorf("DecrementLikesCount: %v", err)
	}

	return nil
}

// Same as IncrementLikesCount, but operates with dislikes count.
func IncrementDislikesCount(database *sql.DB, video_id string) error {
	_, err := database.Exec(
		"INSERT INTO reactions (video_id, likes_count, dislikes_count) VALUES ($1, 0, 1) ON CONFLICT (video_id) DO UPDATE SET dislikes_count = reactions.dislikes_count + 1",
		video_id,
	)

	if err != nil {
		return fmt.Errorf("IncrementDislikesCount: %v", err)
	}

	return nil
}

// Same as DecrementLikesCount, but operates with dislikes count.
func DecrementDislikesCount(database *sql.DB, video_id string) error {
	_, err := database.Exec(
		"INSERT INTO reactions (video_id, likes_count, dislikes_count) VALUES ($1, 0, 0) ON CONFLICT (video_id) DO UPDATE SET dislikes_count = reactions.dislikes_count - 1",
		video_id,
	)

	if err != nil {
		return fmt.Errorf("DecrementDislikesCount: %v", err)
	}

	return nil
}
