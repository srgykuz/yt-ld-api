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
	handlerArgs := createHandlerArgs{
		database: database,
	}
	handler := createHandler(handlerArgs)

	logger.Info(
		fmt.Sprintf("Server is listening on: http://%s", addr),
	)

	err = http.ListenAndServe(addr, handler)

	return err
}

type createHandlerArgs struct {
	database *sql.DB
}

func createHandler(args createHandlerArgs) http.Handler {
	const prefix = "/v0"
	const (
		pathLike          = prefix + "/like"
		pathDislike       = prefix + "/dislike"
		pathRemoveLike    = prefix + "/remove-like"
		pathRemoveDislike = prefix + "/remove-dislike"
		pathStat          = prefix + "/stat"
		pathSignUp        = prefix + "/sign-up"
	)

	mux := http.NewServeMux()

	mux.HandleFunc(
		pathLike,
		wrapCustomHandleFunc(handler.HandleLike, args),
	)
	mux.HandleFunc(
		pathDislike,
		wrapCustomHandleFunc(handler.HandleDislike, args),
	)
	mux.HandleFunc(
		pathRemoveLike,
		wrapCustomHandleFunc(handler.HandleRemoveLike, args),
	)
	mux.HandleFunc(
		pathRemoveDislike,
		wrapCustomHandleFunc(handler.HandleRemoveDislike, args),
	)
	mux.HandleFunc(
		pathStat,
		wrapCustomHandleFunc(handler.HandleStat, args),
	)
	mux.HandleFunc(
		pathSignUp,
		wrapCustomHandleFunc(handler.HandleSignUp, args),
	)

	handler := logReqResMiddleware(mux)

	return handler
}

type customHandleFunc func(args handler.HandlerArgs)

func wrapCustomHandleFunc(f customHandleFunc, args createHandlerArgs) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hA := handler.HandlerArgs{
			Req:      req,
			W:        w,
			Database: args.database,
		}
		f(hA)
	}
}
