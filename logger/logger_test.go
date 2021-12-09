package logger_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Amaimersion/yt-ld-api/logger"
)

func TestInfo(t *testing.T) {
	var out bytes.Buffer
	message := "test"

	logger.SetInfoOutput(&out)
	logger.Info(message)

	result := out.String()
	written := strings.Contains(result, message)

	if !written {
		t.Errorf("result = %s, not contains = %s", result, message)
	}
}

func TestDebug(t *testing.T) {
	var out bytes.Buffer
	message := "test"

	logger.SetDebugOutput(&out)
	logger.Debug(message)

	result := out.String()
	written := strings.Contains(result, message)

	if !written {
		t.Errorf("result = %s, not contains = %s", result, message)
	}
}

func TestOpenLogFileAndWrite(t *testing.T) {
	tmpFile, err := ioutil.TempFile(".", "temp")

	if err != nil {
		t.Fatalf("unable to create temp file: %s", err)
	}

	defer os.Remove(tmpFile.Name())

	// need to close it before opening again
	err = tmpFile.Close()

	if err != nil {
		t.Fatalf("unable to close temp file: %s", err)
	}

	writer, close, err := logger.OpenLogFile(tmpFile.Name())

	if err != nil {
		t.Fatalf("unable to open log file: %s", err)
	}

	_, err = writer.Write([]byte("test"))

	if err != nil {
		t.Fatalf("unable to write in log file: %s", err)
	}

	close()

	// Log file (tmp file) already was closed manually.
	// If for some reason it wasn't, err will be equal to nil.
	err = tmpFile.Close()

	if err == nil {
		t.Error("log file wasn't closed manually")
	}
}
