package handler

import (
	"net/http"
)

// HandleStat handles "statistics" request.
//
// GET will read video statistics (number of likes/dislikes,
// user has liked/disliked the video, etc).
func HandleStat(w http.ResponseWriter, req *http.Request) {
	resp := response{
		status: http.StatusOK,
	}

	switch req.Method {
	case http.MethodGet:
		var args videoInfoArgs
		err := decodeRequestQuery(req, &args)

		if err != nil {
			resp.status = http.StatusBadRequest
			break
		}

		result, err := readStat(args)

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

type readStatResult struct {
	VideoID string `json:"videoID"`
}

func readStat(args videoInfoArgs) (readStatResult, error) {
	result := readStatResult(args)

	return result, nil
}
