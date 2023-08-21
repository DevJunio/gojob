package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	logger  *Logger
	logName string
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func SetLogger(p string) *Logger {
	logger = newLogger(p)
	return logger
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

func (l *Logger) SetOutput(w io.Writer) {
	l.debug.SetOutput(w)
	l.info.SetOutput(w)
	l.warning.SetOutput(w)
	l.err.SetOutput(w)
}

func (l *Logger) Debug(message any) {
	l.debug.Printf(fmt.Sprintf("[DEBUG]: %s\n", message))
}
func (l *Logger) Info(message any) {
	l.info.Printf(fmt.Sprintf("[INFO]: %v\n", message))
}
func (l *Logger) Warn(message any) {
	l.warning.Printf(fmt.Sprintf("[WARN]: %s\n", message))
}
func (l *Logger) Error(message any) {
	l.err.Print(fmt.Sprintf("[ERROR]: %s\n", message))
}

func (l *Logger) Debugf(format string, message ...any) {
	l.debug.Printf(format, message...)
}
func (l *Logger) Warnf(format string, message ...any) {
	l.warning.Printf(format, message...)
}
func (l *Logger) Infof(format string, message ...any) {
	l.info.Printf(format, message...)
}
func (l *Logger) Errorf(format string, message ...any) {
	l.err.Printf(format, message...)
}
