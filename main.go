package main

import (
	"fmt"
	"os"

	"github.com/Amaimersion/yt-alt-ld-api/config"
	"github.com/Amaimersion/yt-alt-ld-api/logger"
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

	defer closeLogs()
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
