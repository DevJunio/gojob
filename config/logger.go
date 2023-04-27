package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func setLog(p string) *log.Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return log.New(writer, p, logger.Flags())
}

func newLogger() *Logger {
	return &Logger{
		debug:   setLog("DEBUG: "),
		info:    setLog("INFO: "),
		warning: setLog("WARNING: "),
		err:     setLog("ERR: "),
		writer:  io.Writer(os.Stdout),
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Printf("debug message: %s\n", v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Printf("info message: %s\n", v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Printf("warn message: %s\n", v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Printf("error message: %s\n", v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
