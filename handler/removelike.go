package handler

import (
	"net/http"
)

// HandleRemoveLike handles "remove like" request.
//
// POST will remove user like on video.
func HandleRemoveLike(w http.ResponseWriter, req *http.Request) {
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

		result, err := deleteLike(args)

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

type deleteLikeResult struct {
	VideoID string `json:"videoID"`
}

func deleteLike(args videoInfoArgs) (deleteLikeResult, error) {
	result := deleteLikeResult(args)

	return result, nil
}
