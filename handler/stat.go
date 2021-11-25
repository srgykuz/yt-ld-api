package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// HandleStat handles "statistics" request.
//
// GET will read video statistics (number of likes/dislikes,
// user has liked/disliked the video, etc).
func HandleStat(w http.ResponseWriter, req *http.Request, database *sql.DB) {
	resp := response{
		status: http.StatusOK,
	}

	switch req.Method {
	case http.MethodGet:
		var args videoInfoArgs

		if err := decodeRequestQuery(req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		result, err := getStat(database, args)

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

	resp.write(w)
}

type getStatResult struct {
	LikesCount    int `json:"likesCount"`
	DislikesCount int `json:"dislikesCount"`
}

func getStat(database *sql.DB, args videoInfoArgs) (getStatResult, error) {
	var result getStatResult
	reaction, err := db.ReadReaction(database, args.VideoID)

	if err != nil {
		return result, err
	}

	result.LikesCount = reaction.LikesCount
	result.DislikesCount = reaction.DislikesCount

	return result, nil
}
