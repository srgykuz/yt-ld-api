package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// encodeResponseBody encodes response Go value into JSON body.
//
// Error will be returned in case if unable to encode properly.
//
// You should pass pointer to value.
func encodeResponseBody(w http.ResponseWriter, v interface{}) error {
	err := json.NewEncoder(w).Encode(v)

	return err
}

// response is a common structure of all responses that will
// be produced by request handlers (except of Go built-in handlers).
type response struct {
	// status is a HTTP code that indicates status of response.
	status int

	// Result is an output of request handler that will be
	// returned to client. Can be of any type that is JSON serializable.
	Result interface{} `json:"result"`
}

func (r *response) write(w http.ResponseWriter) {
	if r.Result != nil {
		w.Header().Set("Content-Type", "application/json")

		if err := encodeResponseBody(w, r); err != nil {
			logger.Info("unable to encode response body: " + err.Error())
		}
	}

	w.WriteHeader(r.status)
}
