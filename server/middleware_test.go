package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogReqResMiddleware(t *testing.T) {
	wantRedirectPath := "/redirect"
	wantStatusCode := http.StatusMovedPermanently

	redirH := http.RedirectHandler(wantRedirectPath, wantStatusCode)
	middH := logReqResMiddleware(redirH)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	middH.ServeHTTP(w, req)

	if w.Code != wantStatusCode {
		t.Errorf("status code = %v, want = %v", w.Code, wantStatusCode)
	}

	if l := w.Header().Get("Location"); l != wantRedirectPath {
		t.Errorf("redirect location = %s, want = %s", l, wantRedirectPath)
	}
}

func TestHandleOptionsMiddleware(t *testing.T) {
	wantRedirectPath := "/redirect"
	wantStatusCode := http.StatusMovedPermanently

	redirH := http.RedirectHandler(wantRedirectPath, wantStatusCode)
	middH := handleOptionsMiddleware(redirH)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	middH.ServeHTTP(w, req)

	if w.Code != wantStatusCode {
		t.Errorf("status code = %v, want = %v", w.Code, wantStatusCode)
	}

	if l := w.Header().Get("Location"); l != wantRedirectPath {
		t.Errorf("redirect location = %s, want = %s", l, wantRedirectPath)
	}
}
