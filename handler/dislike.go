package handler

import (
	"database/sql"
	"net/http"
)

// HandleDislike handles "dislike" request.
//
// POST will mark video as disliked by user.
func HandleDislike(w http.ResponseWriter, req *http.Request, database *sql.DB) {
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

		result, err := createDislike(args)

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

type createDislikeResult struct {
	VideoID string `json:"videoID"`
}

func createDislike(args videoInfoArgs) (createDislikeResult, error) {
	result := createDislikeResult(args)

	return result, nil
}
