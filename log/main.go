package log

import (
	"log"
	"os"
)

var (
	logRaw = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

const (
	logError        = 0
	logInfo         = 1
	logDebug        = 2
	logDetail       = 3
	currentLogLevel = logInfo
)

// Logger abstração para gerar logs
type Logger struct {
	logger      *log.Logger
	application string
	level       int
}

// NewLogger Generate Logger object
func NewLogger(applicationName string) Logger {
	if applicationName == "" {
		applicationName = "UNKNOWN"
	}
	l := Logger{logRaw, applicationName, currentLogLevel}
	return l
}

func (l Logger) Error(e string) {
	l.logger.Printf("[%s]: ERROR: %s", l.application, e)
}

// Info Generate a INFO log line
func (l Logger) Info(text string) {
	if l.level >= logInfo {
		l.logger.Printf("[%s]: INFO: %s", l.application, text)
	}
}

// Debug Generate a DEBUG log line
func (l Logger) Debug(text string) {
	if l.level >= logDebug {
		l.logger.Printf("[%s]: DEBUG: %s", l.application, text)
	}
}

// Panic Generate a PANIC log line and exit
func (l Logger) Panic(err error) {
	l.logger.Panic(err)
}
