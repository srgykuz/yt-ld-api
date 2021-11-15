package server

import (
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"

	"github.com/Amaimersion/yt-alt-ld-api/logger"
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
