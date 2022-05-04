package slogger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	prefix 		string
	output		*io.Writer
}

func (logger Logger) Info(msg string) error {
	return logger.log(mkLogString(Info, msg, logger.prefix))
}

func (logger Logger) Debug(msg string) error {
	return logger.log(mkLogString(Debug, msg, logger.prefix))
}

func (logger Logger) Warn(msg string) error {
	return logger.log(mkLogString(Warn, msg, logger.prefix))
}

func (logger Logger) Err(msg string) error {
	return logger.log(mkLogString(Err, msg, logger.prefix))
}


func (logger Logger) Fatal(msg any) {
	_ = logger.Err(fmt.Sprint(msg))
	os.Exit(1)
}

func (logger Logger) InfoF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return logger.Info(message)
}

func (logger Logger) DebugF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return logger.Info(message)
}

func (logger Logger) WarnF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return logger.Warn(message)
}

func (logger Logger) ErrorF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return logger.Err(message)
}

func (logger Logger) FatalF(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	logger.Fatal(message)
}

func (logger Logger) getLogMessage(msg string, level LogLevel) string {
	return mkLogString(level, msg, logger.prefix)
}

func GetLogger(prefix string) Logger {
	return Logger{output: defaultOutputStream, prefix: prefix}
}

func (logger Logger) log(msg string) error {
	out := logger.output
	_, err := fmt.Fprintln(*out, msg)
	return err
}

func getNowString() string {
	return time.Now().Format(loggingTimeFormat)
}

func mkLogString(level LogLevel, msg string, prefix string) string {
	return fmt.Sprintf("%v [%v] %v: %v", getNowString(), prefix, level, msg)
}

