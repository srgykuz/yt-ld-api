package handler

import (
	"database/sql"
	"net/http"
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
		err := decodeRequestBody(req, &args)

		if err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		result, err := createLike(args)

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

type createLikeResult struct {
	VideoID string `json:"videoID"`
}

func createLike(args videoInfoArgs) (createLikeResult, error) {
	result := createLikeResult(args)

	return result, nil
}
