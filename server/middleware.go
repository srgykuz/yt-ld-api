package server

import (
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"

	"github.com/Amaimersion/yt-ld-api/logger"
)

func logReqResMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)
		s := fmt.Sprintf(
			"%v %v %v %v %v",
			r.Method,
			r.URL.Path,
			m.Code,
			m.Written,
			m.Duration.Milliseconds(),
		)

		logger.Info(s)
	}

	return http.HandlerFunc(fn)
}

func enableCORSMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, Keep-Alive")
		w.Header().Add("Access-Control-Max-Age", "3600")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func handleOptionsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}
