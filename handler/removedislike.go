package handler

import (
	"database/sql"
	"net/http"
)

// HandleRemoveDislike handles "remove dislike" request.
//
// POST will remove user dislike on video.
func HandleRemoveDislike(w http.ResponseWriter, req *http.Request, database *sql.DB) {
	resp := response{
		status: http.StatusOK,
	}

	switch req.Method {
	case http.MethodPost:
		var args videoInfoArgs
		err := decodeRequestBody(req, &args)

		if err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		result, err := deleteDislike(args)

		if err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		resp.Result = result
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(w)
}

type deleteDislikeResult struct {
	VideoID string `json:"videoID"`
}

func deleteDislike(args videoInfoArgs) (deleteDislikeResult, error) {
	result := deleteDislikeResult(args)

	return result, nil
}
