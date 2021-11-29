package db

import (
	"database/sql"
	"fmt"
)

// UserReactions represents all user reactions on specific video.
// Single user can have reactions on many videos, so you always
// need to specify both UserID and VideoID.
type UserReactions struct {
	UserID     int
	VideoID    string
	HasLike    bool
	HasDislike bool
}

// CreateUserReactions creates user reactions with default values.
// Default value for all reactions is false.
func CreateUserReactions(database *sql.DB, userID int, videoID string) error {
	_, err := database.Exec(
		"INSERT INTO user_reactions (user_id, video_id) VALUES ($1, $2)",
		userID,
		videoID,
	)

	if err != nil {
		return fmt.Errorf("CreateUserReactions: %v", err)
	}

	return nil
}

// ReadUserReactions reads all user reactions for specific video.
// If combination of userID and videoID doesn't exists, then ErrNoRow
// will be returned.
func ReadUserReactions(database *sql.DB, userID int, videoID string) (UserReactions, error) {
	result := UserReactions{}
	row := database.QueryRow(
		"SELECT * from user_reactions WHERE user_id = $1 AND video_id = $2",
		userID,
		videoID,
	)
	err := row.Scan(&result.UserID, &result.VideoID, &result.HasLike, &result.HasDislike)

	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRow
		}

		return result, fmt.Errorf("ReadUserReactions: %v", err)
	}

	return result, nil
}

// UpdateUserReactions updates existing user reactions with new value.
// All fields will be updated, so you need to specify new value for every field.
// If combination of userID and videoID doesn't exists, no error will be returned.
func UpdateUserReactions(database *sql.DB, value UserReactions) error {
	_, err := database.Exec(
		"UPDATE user_reactions SET has_like = $1, has_dislike = $2 WHERE user_id = $3 AND video_id = $4",
		value.HasLike,
		value.HasDislike,
		value.UserID,
		value.VideoID,
	)

	if err != nil {
		return fmt.Errorf("UpdateUserReactions: %v", err)
	}

	return nil
}
