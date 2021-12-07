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

	handlerArgs := createHandlerArgs{
		database: database,
		secret:   envCfg.SecretKey,
	}
	handler := createHandler(handlerArgs)
	addr := net.JoinHostPort(host, port)

	logger.Info(
		fmt.Sprintf("Server is listening on: http://%s", addr),
	)

	err = http.ListenAndServe(addr, handler)

	return err
}

type createHandlerArgs struct {
	database *sql.DB
	secret   string
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
		wrapCustomHandlerFunc(handler.HandleLike, args),
	)
	mux.HandleFunc(
		pathDislike,
		wrapCustomHandlerFunc(handler.HandleDislike, args),
	)
	mux.HandleFunc(
		pathRemoveLike,
		wrapCustomHandlerFunc(handler.HandleRemoveLike, args),
	)
	mux.HandleFunc(
		pathRemoveDislike,
		wrapCustomHandlerFunc(handler.HandleRemoveDislike, args),
	)
	mux.HandleFunc(
		pathStat,
		wrapCustomHandlerFunc(handler.HandleStat, args),
	)
	mux.HandleFunc(
		pathSignUp,
		wrapCustomHandlerFunc(handler.HandleSignUp, args),
	)

	var handler http.Handler = mux

	// Order is matters!
	handler = handleOptionsMiddleware(handler)
	handler = logReqResMiddleware(handler)

	return handler
}

type customHandlerFunc func(args handler.HandlerArgs)

func wrapCustomHandlerFunc(f customHandlerFunc, args createHandlerArgs) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hA := handler.HandlerArgs{
			Req:      req,
			W:        w,
			Database: args.database,
			Secret:   args.secret,
		}
		f(hA)
	}
}
