package handler

import (
	"database/sql"
	"net/http"

	"github.com/Amaimersion/yt-alt-ld-api/db"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
)

// HandleSignUp handles "sign up" request.
//
// POST method will sign up the new user.
func HandleSignUp(w http.ResponseWriter, req *http.Request, database *sql.DB) {
	resp := response{
		status: http.StatusOK,
	}

	switch req.Method {
	case http.MethodPost:
		err := signUpUser(database)

		if err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(w)
}

func signUpUser(database *sql.DB) error {
	_, err := db.CreateUser(database)

	if err != nil {
		return err
	}

	return nil
}
