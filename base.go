package slogger

import "io"

// interface of logger
type iLogger interface {
	log(msg string) error
}


type LogLevel string

const (
	Info  LogLevel =	"INFO"
	Debug LogLevel =  "DEBUG"
	Warn  LogLevel =  "WARN"
	Err   LogLevel =  "ERROR"
)

var defaultOutputStream *io.Writer

var loggingTimeFormat string

