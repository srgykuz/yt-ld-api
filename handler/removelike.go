package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// HandleRemoveLike handles "remove like" request.
//
// POST will remove user like on video.
func HandleRemoveLike(w http.ResponseWriter, req *http.Request, database *sql.DB) {
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

		if err := removeLike(database, args); err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(w)
}

func removeLike(database *sql.DB, args videoInfoArgs) error {
	if err := db.DecrementLikesCount(database, args.VideoID); err != nil {
		return err
	}

	return nil
}
