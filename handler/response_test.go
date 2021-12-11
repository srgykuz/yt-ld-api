package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestResponseWrite(t *testing.T) {
	type responseBody struct {
		Key string `json:"key"`
	}

	w := httptest.NewRecorder()
	body := responseBody{
		Key: "value",
	}
	r := response{
		status: 201,
		Result: body,
	}

	r.write(w)

	result := w.Result()

	if result.StatusCode != r.status {
		t.Errorf("status code = %v, want = %v", result.StatusCode, r.status)
	}

	if h := result.Header.Get("Content-Type"); h != "application/json" {
		t.Errorf("Content-Type = %v, want = application/json", h)
	}

	var resultBody struct {
		Result responseBody `json:"result"`
	}

	if err := json.NewDecoder(result.Body).Decode(&resultBody); err == nil {
		if resultBody.Result != body {
			t.Errorf("result body = %v, want = %v", resultBody, body)
		}
	} else {
		t.Errorf("body decode: %v", err)
	}
}
