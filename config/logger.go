package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

var logName string

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func setLog(logInfo string) *log.Logger {
	msg := fmt.Sprint(logName, logInfo)
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, msg, log.Ldate|log.Ltime)

	return log.New(writer, msg, logger.Flags())
}

func newLogger(baseName string) *Logger {
	logName = baseName
	return &Logger{
		debug:   setLog("DEBUG: "),
		info:    setLog("INFO: "),
		warning: setLog("WARNING: "),
		err:     setLog("ERR: "),
		writer:  io.Writer(os.Stdout),
	}
}

func (l *Logger) Debug(message ...interface{}) {
	l.debug.Printf("debug message: %s\n", message...)
}
func (l *Logger) Info(message ...interface{}) {
	l.info.Printf("info message: %s\n", message...)
}
func (l *Logger) Warn(message ...interface{}) {
	l.warning.Printf("warn message: %s\n", message...)
}
func (l *Logger) Error(message ...interface{}) {
	l.err.Printf("error message: %s\n", message...)
}

func (l *Logger) Debugf(format string, message ...interface{}) {
	l.debug.Printf(format, message...)
}
func (l *Logger) Warnf(format string, message ...interface{}) {
	l.warning.Printf(format, message...)
}
func (l *Logger) Infof(format string, message ...interface{}) {
	l.info.Printf(format, message...)
}
func (l *Logger) Errorf(format string, message ...interface{}) {
	l.err.Printf(format, message...)
}
