package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestDecodeRequestBody(t *testing.T) {
	body := videoInfoArgs{
		VideoID: "test",
	}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(&body)

	req := httptest.NewRequest("POST", "/", &b)

	var result videoInfoArgs

	if err := decodeRequestBody(req, &result); err != nil {
		t.Fatalf("decodeRequestBody: %v", err)
	}

	if body != result {
		t.Errorf("result = %v, want = %v", result, body)
	}
}

func TestDecodeInvalidRequestBody(t *testing.T) {
	body := videoInfoArgs{
		VideoID: "",
	}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(&body)

	req := httptest.NewRequest("POST", "/", &b)

	var result videoInfoArgs
	err := decodeRequestBody(req, &result)

	if err != errInvalidRequestBody {
		t.Errorf("err = %v, want = %v", err, errInvalidRequestBody)
	}
}

func TestDecodeRequestQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/?videoID=test", nil)
	var result videoInfoArgs

	if err := decodeRequestQuery(req, &result); err != nil {
		t.Fatalf("err = %v, want = nil", err)
	}

	if result.VideoID != "test" {
		t.Errorf("videoID = %s, want = test", result.VideoID)
	}
}
