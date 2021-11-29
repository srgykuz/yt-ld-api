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
		_, err := parseTokenFromRequest(hArgs.Req, hArgs.Secret)

		if err != nil {
			resp.status = http.StatusUnauthorized
			break
		}

		var args videoInfoArgs

		if err := decodeRequestBody(hArgs.Req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		if err := setLike(hArgs.Database, args); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

func setLike(database *sql.DB, args videoInfoArgs) error {
	if err := db.IncrementLikesCount(database, args.VideoID); err != nil {
		return err
	}

	return nil
}
