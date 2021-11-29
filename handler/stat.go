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
func HandleStat(hArgs HandlerArgs) {
	resp := response{
		status: http.StatusOK,
	}

	switch hArgs.Req.Method {
	case http.MethodGet:
		var args videoInfoArgs

		if err := decodeRequestQuery(hArgs.Req, &args); err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		result, err := getStat(hArgs.Database, args)

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
