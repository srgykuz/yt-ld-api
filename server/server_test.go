package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestEndpointsAvailability(t *testing.T) {
	h := createHandler()
	server := httptest.NewServer(h)

	defer server.Close()

	type endpoint struct {
		path   string
		method string
		name   string
	}

	const wantPrefix = "/v0"
	var wantEndpoints = []endpoint{
		{
			path:   wantPrefix + "/like",
			method: http.MethodPost,
			name:   "POST /like",
		},
		{
			path:   wantPrefix + "/dislike",
			method: http.MethodPost,
			name:   "POST /dislike",
		},
		{
			path:   wantPrefix + "/remove-like",
			method: http.MethodPost,
			name:   "POST /remove-like",
		},
		{
			path:   wantPrefix + "/remove-dislike",
			method: http.MethodPost,
			name:   "POST /remove-dislike",
		},
		{
			path:   wantPrefix + "/stat",
			method: http.MethodGet,
			name:   "GET /stat",
		},
	}

	for _, wantEndpoint := range wantEndpoints {
		t.Run(wantEndpoint.name, func(runT *testing.T) {
			url, err := url.Parse(server.URL + wantEndpoint.path)

			if err != nil {
				runT.Fatal(err)
			}

			req := http.Request{
				Method: wantEndpoint.method,
				URL:    url,
			}
			client := http.Client{}
			res, err := client.Do(&req)

			if err != nil {
				runT.Fatalf("err = %v, want = nil", err)
			}

			if res.StatusCode == http.StatusNotFound {
				runT.Errorf("not found %s", wantEndpoint.path)
			}
		})
	}
}
