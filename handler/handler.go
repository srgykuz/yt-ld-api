package handler

import (
	"database/sql"
	"net/http"
)

// HandlerArgs is an arguments that handler should receive to handle request.
type HandlerArgs struct {
	W        http.ResponseWriter
	Req      *http.Request
	Database *sql.DB
}
