package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type requestBody interface {
	// isValid checks if all needed fields have valid values.
	isValid() bool
}

var (
	errInvalidRequestBody  = errors.New("request body is invalid")
	errInvalidRequestQuery = errors.New("request URL query is invalid")
)

// decodeRequestBody decodes request JSON body into Go value.
// After decoding, value will be checked if all needed fields
// have valid (expected) values.
//
// You should pass pointer to value.
//
// Error will be returned if unable to decode properly.
// errInvalidRequestBody will be returned if result value is invalid.
func decodeRequestBody(req *http.Request, v requestBody) error {
	err := json.NewDecoder(req.Body).Decode(v)

	if err != nil {
		return err
	}

	if !v.isValid() {
		return errInvalidRequestBody
	}

	return nil
}

// decodeRequestQuery is similar to decodeRequestBody, but it performs
// decoding over request URL.
//
// You should pass pointer to value.
//
// Error will be returned if unable to decode properly.
// errInvalidRequestQuery will be returned if result value is invalid.
func decodeRequestQuery(req *http.Request, v requestBody) error {
	switch value := v.(type) {
	case *videoInfoArgs:
		q := req.URL.Query()
		value.fromQuery(q)

		if !value.isValid() {
			return errInvalidRequestQuery
		}
	default:
		panic("unknown type")
	}

	return nil
}

type videoInfoArgs struct {
	VideoID string `json:"videoID"`
}

func (args videoInfoArgs) isValid() bool {
	valid :=
		len(args.VideoID) > 0

	return valid
}

func (args *videoInfoArgs) fromQuery(q url.Values) {
	args.VideoID = q.Get("videoID")
}
