package server

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"github.com/Amaimersion/yt-ld-api/config"
	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/handler"
	"github.com/Amaimersion/yt-ld-api/logger"
)

// ListenAndServe initializes all dependencies that are required to
// serve HTTP requests (DB connection, etc), binds server on given host
// and port. On success it starts listen and serve HTTP requests.
//
// It is a blocking function. This function always returns non-nil error.
func ListenAndServe(host, port string, env config.EnvConfig) error {
	database, err := db.Open(env)

	if err != nil {
		return err
	}

	wrapArgs := wrapHandlerFuncArgs{
		database: database,
		secret:   env.SecretKey,
	}
	handler := createHandler(wrapArgs)
	addr := net.JoinHostPort(host, port)

	logger.Info(fmt.Sprintf("Server is listening on: http://%s", addr))

	err = http.ListenAndServe(addr, handler)

	return err
}

func createHandler(args wrapHandlerFuncArgs) http.Handler {
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
		wrapHandlerFunc(handler.HandleLike, args),
	)
	mux.HandleFunc(
		pathDislike,
		wrapHandlerFunc(handler.HandleDislike, args),
	)
	mux.HandleFunc(
		pathRemoveLike,
		wrapHandlerFunc(handler.HandleRemoveLike, args),
	)
	mux.HandleFunc(
		pathRemoveDislike,
		wrapHandlerFunc(handler.HandleRemoveDislike, args),
	)
	mux.HandleFunc(
		pathStat,
		wrapHandlerFunc(handler.HandleStat, args),
	)
	mux.HandleFunc(
		pathSignUp,
		wrapHandlerFunc(handler.HandleSignUp, args),
	)

	var handler http.Handler = mux

	// Order is matters!
	handler = handleOptionsMiddleware(handler)
	handler = logReqResMiddleware(handler)

	return handler
}

type handlerFunc func(args handler.HandlerArgs)

type wrapHandlerFuncArgs struct {
	database *sql.DB
	secret   string
}

func wrapHandlerFunc(f handlerFunc, args wrapHandlerFuncArgs) http.HandlerFunc {
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
