package logger

import (
	"sync"

	filename "github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var once sync.Once
var logger *log.Logger

func init() {
	// Sigleton
	once.Do(func() {
		logger = log.New()
		formatter := new(prefixed.TextFormatter)
		formatter.FullTimestamp = true

		// Set specific colors for prefix and timestamp
		formatter.SetColorScheme(&prefixed.ColorScheme{
			PrefixStyle:    "cyan+bh",
			TimestampStyle: "black+b:168",
		})
		logger.SetFormatter(formatter)
		logger.AddHook(filename.NewHook())
		logger.SetLevel(log.DebugLevel)
	})
}

func Debug(args ...interface{}) {
	logger.Debugln(args...)
}

func Debugf(s string, args ...interface{}) {
	logger.Debugf(s, args...)
}

func Info(args ...interface{}) {
	logger.Infoln(args...)
}

func Infof(s string, args ...interface{}) {
	logger.Infof(s, args...)
}

func Warning(args ...interface{}) {
	logger.Warnln(args...)
}

func Warningf(s string, args ...interface{}) {
	logger.Warnf(s, args...)
}

func Error(args ...interface{}) {
	logger.Errorln(args...)
}

func Errorf(s string, args ...interface{}) {
	logger.Errorf(s, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatalln(args...)
}

func Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

func Panic(args ...interface{}) {
	logger.Panicln(args...)
}

func Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

// func GetLogger() *log.Logger {
// 	once.Do(func() {
// 		logger = log.New()
// 		formatter := new(prefixed.TextFormatter)
// 		formatter.FullTimestamp = true

// 		// Set specific colors for prefix and timestamp
// 		formatter.SetColorScheme(&prefixed.ColorScheme{
// 			PrefixStyle:    "cyan+bh",
// 			TimestampStyle: "black+b:168",
// 		})
// 		logger.SetFormatter(formatter)
// 		logger.AddHook(filename.NewHook())
// 		logger.SetLevel(log.DebugLevel)
// 	})
// 	return logger
// }
