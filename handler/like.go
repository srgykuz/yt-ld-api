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
func HandleLike(w http.ResponseWriter, req *http.Request, database *sql.DB) {
	resp := response{
		status: http.StatusOK,
	}

	switch req.Method {
	case http.MethodPost:
		var args videoInfoArgs

		if err := decodeRequestBody(req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		if err := setLike(database, args); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(w)
}

func setLike(database *sql.DB, args videoInfoArgs) error {
	if err := db.IncrementLikesCount(database, args.VideoID); err != nil {
		return err
	}

	return nil
}
