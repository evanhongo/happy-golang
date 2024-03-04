package logger

import (
	"io"
	"os"
	"time"

	"github.com/evanhongo/happy-golang/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level
type Field = zap.Field
type Logger struct {
	l     *zap.SugaredLogger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

func (l *Logger) Debug(args ...interface{}) {
	l.l.Debug(args...)
}

func (l *Logger) Debugf(s string, args ...interface{}) {
	l.l.Debugf(s, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.l.Info(args...)
}

func (l *Logger) Infof(s string, args ...interface{}) {
	l.l.Infof(s, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.l.Warn(args...)
}

func (l *Logger) Warnf(s string, args ...interface{}) {
	l.l.Warnf(s, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.l.Error(args...)
}

func (l *Logger) Errorf(s string, args ...interface{}) {
	l.l.Errorf(s, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.l.Panic(args...)
}

func (l *Logger) Panicf(s string, args ...interface{}) {
	l.l.Panicf(s, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.l.Fatal(args...)
}

func (l *Logger) Fatalf(s string, args ...interface{}) {
	l.l.Fatalf(s, args...)
}

func New(writer io.Writer) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	logCfg := zap.NewProductionConfig()
	logCfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}
	cfg := config.GetConfig()
	levelMap := map[string]Level{
		"debug": DebugLevel,
		"info":  InfoLevel,
		"warn":  WarnLevel,
		"error": ErrorLevel,
	}
	level := levelMap[cfg.LOG_LEVEL]
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(logCfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)

	logger := &Logger{
		l:     zap.New(core).Sugar(),
		level: level,
	}
	return logger
}

var std = New(os.Stdout)

func Default() *Logger {
	return std
}

var (
	Binary   = zap.Binary
	Bool     = zap.Bool
	Int8     = zap.Int8
	Int8p    = zap.Int8p
	Float64  = zap.Float64
	Float64p = zap.Float64p
	Float32  = zap.Float32
	Float32p = zap.Float32p
	String   = zap.String
	Strings  = zap.Strings

	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
	Infof  = std.Infof
	Warnf  = std.Warnf
	Errorf = std.Errorf
	Panicf = std.Panicf
	Fatalf = std.Fatalf
	Debugf = std.Debugf
)
