package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Amaimersion/yt-alt-ld-api/config"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
	"github.com/Amaimersion/yt-alt-ld-api/server"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read config: %s", err)
		os.Exit(1)
	}

	closeLogs, err := configureLogger(cfg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to configure logger: %s", err)
		os.Exit(1)
	}

	cleanup := func() {
		closeLogs()
	}

	go func() {
		sigs := listenTerminateSignals()
		sig := <-sigs

		fmt.Fprintf(os.Stderr, "signal: %v\n", sig)
		cleanup()
		os.Exit(0)
	}()

	err = server.ListenAndServe(cfg.Address, cfg.Port)

	fmt.Fprintf(os.Stderr, "listen error: %v\n", err)
	cleanup()
	os.Exit(1)
}

func configureLogger(cfg config.Config) (func(), error) {
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
