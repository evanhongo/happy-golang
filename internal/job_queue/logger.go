package job_queue

import (
	"github.com/RichardKnop/logging"
	"github.com/evanhongo/happy-golang/pkg/logger"
)

type DebugLogger struct {
}

func (l *DebugLogger) Print(args ...interface{}) {
	logger.Debug(args...)
}

// Printf ...
func (l *DebugLogger) Printf(s string, args ...interface{}) {
	logger.Debugf(s, args...)
}

// Println ...
func (l *DebugLogger) Println(args ...interface{}) {
	logger.Debug(args...)
}

// Fatal ...
func (l *DebugLogger) Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf ...
func (l *DebugLogger) Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

// Fatalln ...
func (l *DebugLogger) Fatalln(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic ...
func (l *DebugLogger) Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf ...
func (l *DebugLogger) Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

// Panicln ...
func (l *DebugLogger) Panicln(args ...interface{}) {
	logger.Panic(args...)
}

type InfoLogger struct {
}

func (l *InfoLogger) Print(args ...interface{}) {
	logger.Info(args...)
}

// Printf ...
func (l *InfoLogger) Printf(s string, args ...interface{}) {

	logger.Infof(s, args...)
}

// Println ...
func (l *InfoLogger) Println(args ...interface{}) {
	logger.Info(args...)
}

// Fatal ...
func (l *InfoLogger) Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf ...
func (l *InfoLogger) Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

// Fatalln ...
func (l *InfoLogger) Fatalln(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic ...
func (l *InfoLogger) Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf ...
func (l *InfoLogger) Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

// Panicln ...
func (l *InfoLogger) Panicln(args ...interface{}) {
	logger.Panic(args...)
}

type WarningLogger struct {
}

func (l *WarningLogger) Print(args ...interface{}) {
	logger.Warn(args...)
}

// Printf ...
func (l *WarningLogger) Printf(s string, args ...interface{}) {

	logger.Warnf(s, args...)
}

// Println ...
func (l *WarningLogger) Println(args ...interface{}) {
	logger.Warn(args...)
}

// Fatal ...
func (l *WarningLogger) Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf ...
func (l *WarningLogger) Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

// Fatalln ...
func (l *WarningLogger) Fatalln(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic ...
func (l *WarningLogger) Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf ...
func (l *WarningLogger) Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

// Panicln ...
func (l *WarningLogger) Panicln(args ...interface{}) {
	logger.Panic(args...)
}

type ErrorLogger struct {
}

func (l *ErrorLogger) Print(args ...interface{}) {
	logger.Error(args...)
}

// Printf ...
func (l *ErrorLogger) Printf(s string, args ...interface{}) {

	logger.Errorf(s, args...)
}

// Println ...
func (l *ErrorLogger) Println(args ...interface{}) {
	logger.Error(args...)
}

// Fatal ...
func (l *ErrorLogger) Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf ...
func (l *ErrorLogger) Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

// Fatalln ...
func (l *ErrorLogger) Fatalln(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic ...
func (l *ErrorLogger) Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf ...
func (l *ErrorLogger) Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

// Panicln ...
func (l *ErrorLogger) Panicln(args ...interface{}) {
	logger.Panic(args...)
}

type FatalLogger struct {
}

func (l *FatalLogger) Print(args ...interface{}) {
	logger.Fatal(args...)
}

// Printf ...
func (l *FatalLogger) Printf(s string, args ...interface{}) {

	logger.Fatalf(s, args...)
}

// Println ...
func (l *FatalLogger) Println(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatal ...
func (l *FatalLogger) Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf ...
func (l *FatalLogger) Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

// Fatalln ...
func (l *FatalLogger) Fatalln(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic ...
func (l *FatalLogger) Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf ...
func (l *FatalLogger) Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

// Panicln ...
func (l *FatalLogger) Panicln(args ...interface{}) {
	logger.Panic(args...)
}

func NewLogger(level string) logging.LoggerInterface {
	var logger logging.LoggerInterface
	switch level {
	case "debug":
		logger = &DebugLogger{}
	case "info":
		logger = &InfoLogger{}
	case "warning":
		logger = &WarningLogger{}
	case "error":
		logger = &ErrorLogger{}
	case "fatal":
		logger = &FatalLogger{}
	default:
		logger = &DebugLogger{}
	}
	return logger
}
