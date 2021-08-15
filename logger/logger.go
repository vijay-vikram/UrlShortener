package logger

import (
	"fmt"
	"go.uber.org/zap"
)

// Logger ...
type Logger interface {
	Info(tag string, message string, args ...interface{})
	Error(tag string, message string, args ...interface{})
	Fatal(tag string, message string, args ...interface{})
}

type logger struct {
	ZapClient *zap.SugaredLogger
}

// Initialize ...
func Initialize(logLevel int) {
	var zapLogger *zap.Logger
	// TODO: Handle logs more efficiently
	if logLevel >= 5 {
		zapLogger, _ = zap.NewDevelopment()
		defer zapLogger.Sync().Error()
	} else {
		zapLogger, _ = zap.NewProduction()
		defer zapLogger.Sync().Error()
	}

	Handler = &logger{
		ZapClient: zapLogger.Sugar(),
	}
}

//1 to 5 level logs

//Info ...
func (l *logger) Info(tag string, message string, args ...interface{}) {
	if len(args) > 0 {
		l.ZapClient.Info(tag, fmt.Sprintf(message, args))
	} else {
		l.ZapClient.Info(tag, message)
	}
}

//Error ...
func (l *logger) Error(tag string, message string, args ...interface{}) {
	if len(args) > 0 {
		l.ZapClient.Error(tag, fmt.Sprintf(message, args))
	} else {
		l.ZapClient.Error(tag, message)
	}
}

//Fatal ...
func (l *logger) Fatal(tag string, message string, args ...interface{}) {
	if len(args) > 0 {
		l.ZapClient.Fatal(tag, fmt.Sprintf(message, args))
	} else {
		l.ZapClient.Fatal(tag, message)
	}
}
