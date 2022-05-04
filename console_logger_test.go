package slogger


import (
	"io"
	"os"
	"testing"
)
import color "github.com/AminMal/slogger/colored"
import "reflect"

func TestNewConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger("TestCase")
	logErr := logger.log(LogItem{
		color: color.CYAN,
		msg:   "Test case here",
		level: Info,
	})
	reflect.TypeOf(logger.underlying.output)
	var stderr io.Writer = os.Stderr
	isWritingToConsole := reflect.TypeOf(*(logger.underlying.output)) == reflect.TypeOf(stderr)
	hasSamePrefix := logger.underlying.prefix == "TestCase"
	if !isWritingToConsole {
		t.Fatal("logger is not writing to the stderr")
	}
	if !hasSamePrefix {
		t.Fatalf(
			"logger doesn't have the same prefix as underlying logger, consolelogger prefix: %v, underlying logger prefix: %v",
			"TestCase",
			logger.underlying.prefix,
		)
	}
	if logErr != nil {
		t.Fatalf("logger couldn't log as expected, %v", logErr.Error())
	}
}