package config

import (
	"os"

	"github.com/joho/godotenv"
)

// EnvConfig is a configuration that can be obtained
// from environment variables.
type EnvConfig struct {
	// Name of database user.
	DBUser string

	// Password of database user.
	DBPassword string

	// Host of database.
	DBHost string

	// Port of database.
	DBPort string

	// Name of database.
	DBName string
}

// ReadEnv reads configuration from environment variables.
//
// If some key not presented in environment variables, then
// its value will be equal to empty value.
//
// If you want to read environment variables from file,
// then call LoadEnv() before this function.
func ReadEnv() EnvConfig {
	cfg := EnvConfig{}

	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")

	return cfg
}

// LoadEnv reads environment variables from file and loads them.
// You should call this function as soon as possible, certainly before ReadEnv().
func LoadEnv(file string) error {
	return godotenv.Load(file)
}
