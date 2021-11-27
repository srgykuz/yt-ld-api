package handler

import (
	"database/sql"
	"net/http"
)

// HandleSignUp handles "sign up" request.
func HandleSignUp(w http.ResponseWriter, req *http.Request, database *sql.DB) {
	resp := response{
		status: http.StatusOK,
	}

	switch req.Method {
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(w)
}
