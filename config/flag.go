package config

import (
	"flag"
)

// FlagConfig is a configuration that can be obtained
// from command line flags.
type FlagConfig struct {
	// Bind server to this IP address.
	Host string

	// Bind server to this TCP port.
	Port string

	// Output info logs to this file.
	// Use /dev/null to disable info logs.
	InfoOutput string

	// Output debug logs to this file.
	// Use /dev/null to disable debug logs.
	DebugOutput string

	// Load environment variables from this file.
	// Use empty string if nothing should be read.
	EnvFile string
}

const (
	defaultHost        = "0.0.0.0"
	defaultPort        = "8080"
	defaultInfoOutput  = "/dev/stderr"
	defaultDebugOutput = "/dev/null"
	defaultEnvFile     = ""
)

// ReadFlags reads configuration from command line flags.
//
// Default values will be used if some configuration value
// is not presented.
func ReadFlags() FlagConfig {
	cfg := FlagConfig{}

	flag.StringVar(
		&cfg.Host,
		"host",
		defaultHost,
		"Bind server to this host. IP or host name.",
	)
	flag.StringVar(
		&cfg.Port,
		"port",
		defaultPort,
		"Bind server to this TCP port. Number or service name.",
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
	flag.StringVar(
		&cfg.EnvFile,
		"envFile",
		defaultEnvFile,
		"Load environment variables from this file.",
	)

	flag.Parse()

	return cfg
}
