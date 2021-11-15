package server

import (
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"

	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// ListenAndServe executes server on given host and port.
// It will listen and serve HTTP requests.
//
// It is a blocking function. This function always returns non-nil error.
func ListenAndServe(host string, port int) error {
	addr := host + ":" + fmt.Sprint(port)
	handler := createHandler()

	logger.Info(
		fmt.Sprintf("Server is listening on: http://%s", addr),
	)

	err := http.ListenAndServe(addr, handler)

	return err
}

func createHandler() http.Handler {
	mux := http.NewServeMux()

	handler := logReqResMiddleware(mux)

	return handler
}

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
