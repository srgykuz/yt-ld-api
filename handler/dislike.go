package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// HandleDislike handles "dislike" request.
//
// POST will mark video as disliked by user.
func HandleDislike(hArgs HandlerArgs) {
	resp := response{
		status: http.StatusOK,
	}

	switch hArgs.Req.Method {
	case http.MethodPost:
		var args videoInfoArgs

		if err := decodeRequestBody(hArgs.Req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		if err := setDislike(hArgs.Database, args); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

func setDislike(database *sql.DB, args videoInfoArgs) error {
	if err := db.IncrementDislikesCount(database, args.VideoID); err != nil {
		return err
	}

	return nil
}
