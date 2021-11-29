package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// HandleLike handles "like" request.
//
// POST will mark video as liked by user.
func HandleLike(hArgs HandlerArgs) {
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

		if err := setLike(hArgs.Database, args, token.UserID); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

func setLike(database *sql.DB, args videoInfoArgs, userID int) error {
	userReactions, err := db.ReadUserReactions(database, userID, args.VideoID)
	create := false

	if err != nil {
		if err == db.ErrNoRow {
			create = true
		} else {
			return err
		}
	}

	if create {
		if err := db.CreateUserReactions(database, userID, args.VideoID); err != nil {
			return err
		}

		userReactions = db.UserReactions{
			UserID:  userID,
			VideoID: args.VideoID,
		}
	}

	if create || !userReactions.HasLike {
		if userReactions.HasDislike {
			if err := db.DecrementDislikesCount(database, args.VideoID); err != nil {
				return err
			}
		}

		if err := db.IncrementLikesCount(database, args.VideoID); err != nil {
			return err
		}

		userReactions.HasLike = true
		userReactions.HasDislike = false

		if err := db.UpdateUserReactions(database, userReactions); err != nil {
			return err
		}
	}

	return nil
}
