package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// HandleRemoveDislike handles "remove dislike" request.
//
// POST will remove user dislike on video.
func HandleRemoveDislike(hArgs HandlerArgs) {
	resp := response{
		status: http.StatusOK,
	}

	switch hArgs.Req.Method {
	case http.MethodPost:
		token, err := parseTokenFromRequest(hArgs.Req, hArgs.Secret)

		if err != nil {
			resp.status = http.StatusUnauthorized
			break
		}

		var args videoInfoArgs

		if err := decodeRequestBody(hArgs.Req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		if err := removeDislike(hArgs.Database, args, token.UserID); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

func removeDislike(database *sql.DB, args videoInfoArgs, userID int) error {
	if err := db.DecrementDislikesCount(database, args.VideoID); err != nil {
		return err
	}

	userReactions, err := db.ReadUserReactions(database, userID, args.VideoID)

	if err != nil {
		if err == db.ErrNoRow {
			return nil
		}

		return err
	}

	if userReactions.HasDislike {
		userReactions.HasDislike = false

		if err := db.UpdateUserReactions(database, userReactions); err != nil {
			return err
		}
	}

	return nil
}
