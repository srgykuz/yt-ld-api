// Package logger provides logging methods for entire program.
//
// Note that there only two types of logs: info and debug.
// Other types like warning, error, critical, etc.
// are considered as useless and even harmful.
// For example, if you want to handle error, handle it properly or
// pass to the caller; if you want to warn about something,
// then code it properly that no warn logs are needed; if you want
// to exit the program, then clear the program and exit it properly.
//
// Info level is intended for users who launched this program and want to
// monitor it output to understand program state. Debug level is intended
// for developers who debugging the program at the moment.
//
// Don't garbage logs, don't use loggers for everything.
// Always try to do necessary things through code.
// Log information only when it is really useful when reading logs.
package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	infoLogger = log.New(
		os.Stderr,
		"",
		log.LstdFlags,
	)
	debugLogger = log.New(
		os.Stderr,
		"[DEBUG] ",
		log.Ltime|log.Lmicroseconds,
	)
)

// Info writes message at info level.
func Info(s string) {
	infoLogger.Print(s)
}

// Debug writes message at debug level.
func Debug(s string) {
	debugLogger.Print(s)
}

// SetInfoOutput sets output for info logs.
//
// Pass os.DevNull to disable logging.
func SetInfoOutput(w io.Writer) {
	infoLogger.SetOutput(w)
}

// SetDebugOutput sets output for debug logs.
//
// Pass os.DevNull to disable logging.
func SetDebugOutput(w io.Writer) {
	debugLogger.SetOutput(w)
}

// OpenLogFile opens file for logging purposes.
// If file not exists, it will be created.
// If file exists, new content will be appended.
//
// Name is a file path. It also can be /dev/stdout, /dev/stderr,
// /dev/null, /dev/zero. Appropriate writer will be returned
// that can be used for SetInfoOutput() or SetDebugOutput().
//
// Close function will be returned. This function closes opened file.
// You must call this function.
//
// Error will be returned if unable to open file. In that case you
// can skip calling of close function.
func OpenLogFile(name string) (io.Writer, func(), error) {
	closeVoid := func() {}

	switch name {
	case "/dev/stdout":
		return os.Stdout, closeVoid, nil
	case "/dev/stderr":
		return os.Stderr, closeVoid, nil
	case "/dev/null":
	case "/dev/zero":
		return ioutil.Discard, closeVoid, nil
	}

	file, err := os.OpenFile(
		name,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		os.ModePerm,
	)

	if err != nil {
		return nil, closeVoid, err
	}

	closeFile := func() {
		file.Close()
	}

	return file, closeFile, nil
}
