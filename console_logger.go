package slogger

import (
	"fmt"
	"io"
	"os"
	color "github.com/AminMal/slogger/colored"
)

type LogItem struct {
	color 	color.Color
	msg 	string
	level   LogLevel
}

type ConsoleLogger struct {
	underlying 		Logger
}

func NewConsoleLogger(prefix string) ConsoleLogger {
	var stderr io.Writer
	stderr = os.Stderr
	return ConsoleLogger{underlying: Logger{prefix: prefix, output: &stderr}}
}

func (cl ConsoleLogger) log(item LogItem) error {
	loggerMessage := cl.underlying.getLogMessage(item.msg, item.level)
	var finalMessage string
	if item.color != color.NoColor {
		finalMessage = fmt.Sprintf("%v%v%v", item.color, loggerMessage, color.ResetPrevColor)
	} else {
		finalMessage = loggerMessage
	}
	return cl.underlying.log(finalMessage)
}

func (cl ConsoleLogger) Info(msg string) error {
	return cl.log(LogItem{
		color: color.NoColor,
		msg:   msg,
		level: Info,
	})
}

func (cl ConsoleLogger) Debug(msg string) error {
	return cl.log(LogItem{
		color: color.NoColor,
		msg:   msg,
		level: Debug,
	})
}

func (cl ConsoleLogger) Warn(msg string) error {
	return cl.log(LogItem{
		color: color.NoColor,
		msg:   msg,
		level: Warn,
	})
}

func (cl ConsoleLogger) Err(msg string) error {
	return cl.log(LogItem{
		color: color.NoColor,
		msg:   msg,
		level: Err,
	})
}

func (cl ConsoleLogger) Fatal(msg any) {
	_ = cl.Err(fmt.Sprint(msg))
	os.Exit(1)
}

func (cl ConsoleLogger) InfoF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return cl.Info(message)
}

func (cl ConsoleLogger) DebugF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return cl.Debug(message)
}

func (cl ConsoleLogger) WarnF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return cl.Warn(message)
}

func (cl ConsoleLogger) ErrorF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return cl.Err(message)
}

func (cl ConsoleLogger) FatalF(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	cl.Fatal(message)
}



type LogItemBuilder struct {
	color	color.Color
	level   LogLevel
	cl      *ConsoleLogger
}

func (cl ConsoleLogger) Colored(color color.Color) LogItemBuilder {
	return LogItemBuilder{color: color, cl: &cl}
}

func (lib LogItemBuilder) build(msg string) LogItem {
	return LogItem {
		color: lib.color,
		msg: msg,
		level: lib.level,
	}
}

func (lib LogItemBuilder) withLevel(level LogLevel) LogItemBuilder {
	return LogItemBuilder{
		color: lib.color,
		level: level,
		cl:    lib.cl,
	}
}

func (lib LogItemBuilder) Info(msg string) error {
	logItem := lib.withLevel(Info).build(msg)
	return lib.cl.log(logItem)
}

func (lib LogItemBuilder) Debug(msg string) error {
	logItem := lib.withLevel(Debug).build(msg)
	return lib.cl.log(logItem)
}

func (lib LogItemBuilder) Warn(msg string) error {
	logItem := lib.withLevel(Warn).build(msg)
	return lib.cl.log(logItem)
}

func (lib LogItemBuilder) Err(msg string) error {
	logItem := lib.withLevel(Err).build(msg)
	return lib.cl.log(logItem)
}

func (lib LogItemBuilder) Fatal(msg any) {
	_ = lib.Err(fmt.Sprint(msg))
	os.Exit(1)
}

func (lib LogItemBuilder) InfoF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return lib.Info(message)
}

func (lib LogItemBuilder) DebugF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return lib.Debug(message)
}

func (lib LogItemBuilder) WarnF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return lib.Warn(message)
}

func (lib LogItemBuilder) ErrorF(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return lib.Err(message)
}

func (lib LogItemBuilder) FatalF(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	lib.Fatal(message)
}