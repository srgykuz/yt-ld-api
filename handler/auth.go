package handler

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// tokenData is a custom data that will be embed in token.
type tokenData struct {
	// User DB ID.
	UserID int `json:"userID"`
}

type tokenClaims struct {
	tokenData
	jwt.RegisteredClaims
}

// createToken creates JWT token that can be used for
// authorization purposes.
//
// data is a custom data that will be embed in token,
// now is a current time, secret is a constatnt secret key.
//
// In case of error don't return this error to user,
// instead create new one with uninformative message.
func createToken(data tokenData, now time.Time, secret string) (string, error) {
	claims := tokenClaims{
		data,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))

	return ss, err
}

// parseToken parses token that was created with createToken().
func parseToken(tokenString string, secret string) (tokenData, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
	claims := tokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)

	if err != nil {
		return claims.tokenData, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.tokenData, nil
	}

	return claims.tokenData, errors.New("token is not valid")
}
