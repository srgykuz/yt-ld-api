package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/logger"
)

// HandleStat handles "statistics" request.
//
// GET will read video statistics (number of likes/dislikes,
// user has liked/disliked the video, etc).
func HandleStat(hArgs HandlerArgs) {
	resp := response{
		status: http.StatusOK,
	}

	switch hArgs.Req.Method {
	case http.MethodGet:
		token, err := parseTokenFromRequest(hArgs.Req, hArgs.Secret)

		if err != nil {
			resp.status = http.StatusUnauthorized
			break
		}

		var args videoInfoArgs

		if err := decodeRequestQuery(hArgs.Req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		result, err := getStat(hArgs.Database, args, token.UserID)

		if err != nil {
			if err == db.ErrNoRow {
				resp.status = http.StatusNotFound
			} else {
				resp.status = http.StatusInternalServerError
				logger.Info(err.Error())
			}

			break
		}

		resp.Result = result
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

type getStatResult struct {
	LikesCount    int  `json:"likesCount"`
	DislikesCount int  `json:"dislikesCount"`
	HasLike       bool `json:"hasLike"`
	HasDislike    bool `json:"hasDislike"`
}

func getStat(database *sql.DB, args videoInfoArgs, userID int) (getStatResult, error) {
	var result getStatResult

	reaction, err := db.ReadReaction(database, args.VideoID)

	if err != nil {
		if err == db.ErrNoRow {
			result := getStatResult{
				LikesCount:    0,
				DislikesCount: 0,
				HasLike:       false,
				HasDislike:    false,
			}

			return result, nil
		}

		return result, err
	}

	userReactions, err := db.ReadUserReactions(database, userID, args.VideoID)

	if err != nil {
		if err == db.ErrNoRow {
			userReactions = db.UserReactions{
				HasLike:    false,
				HasDislike: false,
			}
		} else {
			return result, err
		}
	}

	result.LikesCount = reaction.LikesCount
	result.DislikesCount = reaction.DislikesCount
	result.HasLike = userReactions.HasLike
	result.HasDislike = userReactions.HasDislike

	return result, nil
}
