package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareNotInterruptsCallChain(t *testing.T) {
	type middlewareTest struct {
		name       string
		middleware func(http.Handler) http.Handler
	}

	tests := []middlewareTest{
		{
			name:       "logReqRes",
			middleware: logReqResMiddleware,
		},
		{
			name:       "enableCORS",
			middleware: enableCORSMiddleware,
		},
		{
			name:       "handleOptions",
			middleware: handleOptionsMiddleware,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(runT *testing.T) {
			wantRedirectPath := "/redirect"
			wantStatusCode := http.StatusMovedPermanently

			redirH := http.RedirectHandler(wantRedirectPath, wantStatusCode)
			middH := test.middleware(redirH)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			middH.ServeHTTP(w, req)

			if w.Code != wantStatusCode {
				runT.Errorf("status code = %v, want = %v", w.Code, wantStatusCode)
			}

			if l := w.Header().Get("Location"); l != wantRedirectPath {
				runT.Errorf("redirect location = %s, want = %s", l, wantRedirectPath)
			}
		})
	}
}
