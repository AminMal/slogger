package slogger

import (
	"io"
	"os"
)

func init() {
	// setting default output stream to console
	// should change this function if you want to log to another output stream
	var stderr io.Writer
	stderr = os.Stderr
	defaultOutputStream = &stderr

	loggingTimeFormat = "2006/01/02 15:04:05"
}
