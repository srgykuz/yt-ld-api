package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/logger"
)

// HandleSignUp handles "sign up" request.
//
// POST method will sign up the new user.
func HandleSignUp(hArgs HandlerArgs) {
	resp := response{
		status: http.StatusOK,
	}

	switch hArgs.Req.Method {
	case http.MethodPost:
		result, err := signUpUser(hArgs.Database, hArgs.Secret)

		if err != nil {
			resp.status = http.StatusInternalServerError
			logger.Info(err.Error())
			break
		}

		resp.Result = result
	default:
		resp.status = http.StatusMethodNotAllowed
	}

	resp.write(hArgs.W)
}

type signUpUserResult struct {
	AccessToken string `json:"accessToken"`
}

func signUpUser(database *sql.DB, secret string) (signUpUserResult, error) {
	result := signUpUserResult{}
	id, err := db.CreateUser(database)

	if err != nil {
		return result, err
	}

	token, err := createToken(
		tokenData{UserID: id},
		time.Now(),
		secret,
	)

	if err != nil {
		return result, err
	}

	result.AccessToken = token

	return result, nil
}
