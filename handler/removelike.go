package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/logger"
)

// HandleRemoveLike handles "remove like" request.
//
// POST will remove user like on video.
func HandleRemoveLike(hArgs HandlerArgs) {
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

		if err := removeLike(hArgs.Database, args, token.UserID); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

func removeLike(database *sql.DB, args videoInfoArgs, userID int) error {
	userReactions, err := db.ReadUserReactions(database, userID, args.VideoID)

	if err != nil {
		if err == db.ErrNoRow {
			return nil
		}

		return err
	}

	if userReactions.HasLike {
		if err := db.DecrementLikesCount(database, args.VideoID); err != nil {
			return err
		}

		userReactions.HasLike = false

		if err := db.UpdateUserReactions(database, userReactions); err != nil {
			return err
		}
	}

	return nil
}
