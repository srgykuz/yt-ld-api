package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/handler"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// ListenAndServe executes server on given host and port.
// It will listen and serve HTTP requests.
//
// It is a blocking function. This function always returns non-nil error.
func ListenAndServe(host, port string) error {
	addr := net.JoinHostPort(host, port)
	handler := createHandler()

	logger.Info(
		fmt.Sprintf("Server is listening on: http://%s", addr),
	)

	err := http.ListenAndServe(addr, handler)

	return err
}

func createHandler() http.Handler {
	const prefix = "/v0"
	const (
		pathLike          = prefix + "/like"
		pathDislike       = prefix + "/dislike"
		pathRemoveLike    = prefix + "/remove-like"
		pathRemoveDislike = prefix + "/remove-dislike"
		pathStat          = prefix + "/stat"
	)

	mux := http.NewServeMux()

	mux.HandleFunc(pathLike, handler.HandleLike)
	mux.HandleFunc(pathDislike, handler.HandleDislike)
	mux.HandleFunc(pathRemoveLike, handler.HandleRemoveLike)
	mux.HandleFunc(pathRemoveDislike, handler.HandleRemoveDislike)
	mux.HandleFunc(pathStat, handler.HandleStat)

	handler := logReqResMiddleware(mux)

	return handler
}
