package slogger

import (
	"log"
	"regexp"
	"testing"
)

func failWith(t *testing.T, message string) {
	t.Fatalf("the following log result: \"%v\" does not conform to the expected expression", message)
}

func TestLogger_Info(t *testing.T) {
	logger := GetLogger("TestCase1")
	finalMessage := logger.getLogMessage("Hi, this is an info log", Info)
	expectedRegex, _ := regexp.Compile(".+\\s\\[TestCase1\\] INFO: Hi, this is an info log")
	if ok := expectedRegex.Match([]byte(finalMessage)); !ok {
		failWith(t, finalMessage)
	}
}

func TestLogger_Debug(t *testing.T) {
	logger := GetLogger("TestCase1")
	finalMessage := logger.getLogMessage("Hi, this is a debug log", Debug)
	expectedRegexp, _ := regexp.Compile(".+\\s\\[TestCase1\\] DEBUG: Hi, this is a debug log")
	if ok := expectedRegexp.Match([]byte(finalMessage)); !ok {
		failWith(t, finalMessage)
	}
}

func TestLogger_Err(t *testing.T) {
	logger := GetLogger("TestCase1")
	finalMessage := logger.getLogMessage("Hi, this is a failure log", Err)
	expectedRegexp, _ := regexp.Compile(".+\\s\\[TestCase1\\] ERROR: Hi, this is a failure log")
	if ok := expectedRegexp.Match([]byte(finalMessage)); !ok {
		failWith(t, finalMessage)
	}
}

func TestLogger_Warn(t *testing.T) {
	logger := GetLogger("TestCase1")
	finalMessage := logger.getLogMessage("Hi, this is a warning message", Warn)
	expectedRegexp, _ := regexp.Compile(".+\\s\\[TestCase1\\] WARN: Hi, this is a warning message")
	if ok := expectedRegexp.Match([]byte(finalMessage)); !ok {
		failWith(t, finalMessage)
	}
}

var benchMarkMessage = "bench mark message to log"

func BenchmarkLogger_Info(b *testing.B) {
	logger := GetLogger("Benchmark")
	for i := 0; i < b.N; i ++ {
		logger.Info(benchMarkMessage)
	}
}

func BenchmarkDefaultLogger(b *testing.B) {
	logger := log.Default()
	logger.SetPrefix("[Benchmark] Info: ")
	for i := 0; i < b.N; i++ {
		logger.Println(benchMarkMessage)
	}
}
