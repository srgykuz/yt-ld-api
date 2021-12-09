package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Amaimersion/yt-ld-api/config"
	"github.com/Amaimersion/yt-ld-api/db"
	"github.com/Amaimersion/yt-ld-api/logger"
	"github.com/Amaimersion/yt-ld-api/server"
)

func main() {
	flags := config.ReadFlags()

	if len(flags.EnvFile) > 0 {
		if err := config.LoadEnv(flags.EnvFile); err != nil {
			fmt.Fprintf(os.Stderr, "LoadEnv: %v\n", err)
			os.Exit(1)
		}
	}

	closeLogs, err := configureLogger(flags)

	if err != nil {
		fmt.Fprintf(os.Stderr, "configureLogger: %v\n", err)
		os.Exit(1)
	}

	cleanup := func() {
		closeLogs()
	}

	go func() {
		sigs := listenTerminateSignals()
		sig := <-sigs

		fmt.Fprintf(os.Stdout, "signal: %v\n", sig)
		cleanup()
		os.Exit(0)
	}()

	env := config.ReadEnv()
	serverArgs := server.ListenAndServeArgs{
		Host:   flags.Host,
		Port:   flags.Port,
		Secret: env.SecretKey,
		Database: db.OpenArgs{
			User:     env.DBUser,
			Password: env.DBPassword,
			Host:     env.DBHost,
			Port:     env.DBPort,
			Name:     env.DBName,
		},
	}
	err = server.ListenAndServe(serverArgs)

	fmt.Fprintf(os.Stderr, "ListenAndServe: %v\n", err)
	cleanup()
	os.Exit(1)
}

// configureLogger configures global app logger.
//
// Use returned function to close log files.
func configureLogger(cfg config.FlagConfig) (func(), error) {
	infoF, closeInfoF, err := logger.OpenLogFile(cfg.InfoOutput)

	if err != nil {
		return func() {}, err
	}

	debugF, closeDebugF, err := logger.OpenLogFile(cfg.DebugOutput)

	if err != nil {
		closeInfoF()
		return func() {}, err
	}

	closeF := func() {
		closeInfoF()
		closeDebugF()
	}

	logger.SetInfoOutput(infoF)
	logger.SetDebugOutput(debugF)

	return closeF, nil
}

func listenTerminateSignals() <-chan os.Signal {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	return sigs
}
