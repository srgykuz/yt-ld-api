package config

import (
	"flag"
)

// Config is the app configuration.
type Config struct {
	// Bind server to this IP.
	Address string

	// Bind server to this TCP port.
	Port int

	// Output info logs to this file.
	// Use /dev/null to disable info logs.
	InfoOutput string

	// Output debug logs to this file.
	// Use /dev/null to disable debug logs.
	DebugOutput string
}

const (
	defaultAddress     = "0.0.0.0"
	defaultPort        = 8080
	defaultInfoOutput  = "/dev/stderr"
	defaultDebugOutput = "/dev/stderr"
)

// Read reads configuration from available place.
//
// Default values will be used if some configuration value
// is not presented. If unable to read at all, then error
// will be returned.
func Read() (Config, error) {
	cfg := readFlags()

	return cfg, nil
}

func readFlags() Config {
	cfg := Config{}

	flag.StringVar(
		&cfg.Address,
		"address",
		defaultAddress,
		"Bind server to this IP.",
	)
	flag.IntVar(
		&cfg.Port,
		"port",
		defaultPort,
		"Bind server to this TCP port.",
	)
	flag.StringVar(
		&cfg.InfoOutput,
		"infoLog",
		defaultInfoOutput,
		"Output info logs to this file.",
	)
	flag.StringVar(
		&cfg.DebugOutput,
		"debugLog",
		defaultDebugOutput,
		"Output debug logs to this file.",
	)

	flag.Parse()

	return cfg
}
