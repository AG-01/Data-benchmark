package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger wraps logrus logger with custom functionality
type Logger struct {
	*logrus.Logger
}

// New creates a new logger instance
func New() *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	
	return &Logger{Logger: log}
}

// WithFields creates a new logger with fields
func (l *Logger) WithFields(fields map[string]interface{}) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

// WithError creates a new logger with error
func (l *Logger) WithError(err error) *logrus.Entry {
	return l.Logger.WithError(err)
}
