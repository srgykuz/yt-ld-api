package server

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/config"
	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/handler"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// ListenAndServe initializes all dependencies that are required to
// serve HTTP requests (DB connection, etc), binds server on given host
// and port. On success it starts listen and serve HTTP requests.
//
// It is a blocking function. This function always returns non-nil error.
func ListenAndServe(host, port string, envCfg config.EnvConfig) error {
	database, err := db.Open(envCfg)

	if err != nil {
		return err
	}

	addr := net.JoinHostPort(host, port)
	handler := createHandler(database)

	logger.Info(
		fmt.Sprintf("Server is listening on: http://%s", addr),
	)

	err = http.ListenAndServe(addr, handler)

	return err
}

func createHandler(database *sql.DB) http.Handler {
	const prefix = "/v0"
	const (
		pathLike          = prefix + "/like"
		pathDislike       = prefix + "/dislike"
		pathRemoveLike    = prefix + "/remove-like"
		pathRemoveDislike = prefix + "/remove-dislike"
		pathStat          = prefix + "/stat"
	)

	mux := http.NewServeMux()

	mux.HandleFunc(
		pathLike,
		wrapCustomHandleFunc(handler.HandleLike, database),
	)
	mux.HandleFunc(
		pathDislike,
		wrapCustomHandleFunc(handler.HandleDislike, database),
	)
	mux.HandleFunc(
		pathRemoveLike,
		wrapCustomHandleFunc(handler.HandleRemoveLike, database),
	)
	mux.HandleFunc(
		pathRemoveDislike,
		wrapCustomHandleFunc(handler.HandleRemoveDislike, database),
	)
	mux.HandleFunc(
		pathStat,
		wrapCustomHandleFunc(handler.HandleStat, database),
	)

	handler := logReqResMiddleware(mux)

	return handler
}

type customHandleFunc func(http.ResponseWriter, *http.Request, *sql.DB)

func wrapCustomHandleFunc(f customHandleFunc, database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		f(w, req, database)
	}
}
