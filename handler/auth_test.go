package handler

import (
	"testing"
	"time"
)

const secret = "super-secret"
const resultToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsIm5iZiI6MTE4ODYzNzIwMCwiaWF0IjoxMTg4NjM3MjAwfQ.-pOIeZl7rXPa5--M5wi4To-M19RQXwPq5W6IXbTZd_E"

var resultData = tokenData{
	UserID: 1,
}
var now = time.Date(2007, time.September, 1, 9, 0, 0, 0, time.UTC)

func TestCreateToken(t *testing.T) {
	token, err := createToken(resultData, now, secret)

	if err != nil {
		t.Fatalf("err = %v, want = nil", err)
	}

	if token != resultToken {
		t.Errorf("token = %v, want = %v", token, resultToken)
	}
}

func TestParseToken(t *testing.T) {
	data, err := parseToken(resultToken, secret)

	if err != nil {
		t.Fatalf("err = %v, want = nil", err)
	}

	if data.UserID != resultData.UserID {
		t.Errorf("UserID = %v, want = %v", data.UserID, resultData.UserID)
	}
}

func TestParseTokenInvalid(t *testing.T) {
	tokenString := resultToken[1:]
	_, err := parseToken(tokenString, secret)

	if err == nil {
		t.Fatal("err = nil, want some error")
	}
}

func TestParseTokenDifferentSecret(t *testing.T) {
	newSecret := secret[1:]
	_, err := parseToken(resultToken, newSecret)

	if err == nil {
		t.Fatal("err = nil, want some error")
	}
}
