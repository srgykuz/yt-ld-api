package db

import (
	"database/sql"
	"fmt"
	"time"
)

// User represents user model.
type User struct {
	ID         int
	SignUpDate time.Time
}

// CreateUser creates new user with default values.
// You don't need to pass any other data because at the
// moment default values describes enough user data.
// On success, ID of created user will be returned.
func CreateUser(database *sql.DB) (int, error) {
	row := database.QueryRow("INSERT INTO users DEFAULT VALUES RETURNING id")
	var id int

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}

	return id, nil
}
