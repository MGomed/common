package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	lumberjack "github.com/natefinch/lumberjack"
	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

const (
	logTimeLayout = "02-01-2006_03:04:05"

	outDir = "out/log"
)

// Interface is Logger and TestLogger interface
type Interface interface {
	Debug(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{})
	Error(format string, params ...interface{})
	Panic(format string, params ...interface{})
}

// Logger is zap.Logger wrapper
type Logger struct {
	log *zap.Logger
}

// NewLogger is Logger struct constructor
func NewLogger(serviceName string) Interface {
	var cores []zapcore.Core
	if params.isConsoleEncoder {
		params.consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(params.consoleConfig),
			zapcore.AddSync(os.Stdout),
			params.logLevel,
		)

		cores = append(cores, consoleCore)
	}

	if params.isFileEncoder {
		params.fileConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(params.fileConfig),
			zapcore.AddSync(&lumberjack.Logger{
				Filename: path.Join(outDir, fmt.Sprintf("%s_%s.log",
					serviceName, time.Now().Format(logTimeLayout))),
				MaxSize: 10,
				MaxAge:  7,
			}),

			params.logLevel,
		)

		cores = append(cores, fileCore)
	}

	return &Logger{
		log: zap.New(
			zapcore.NewTee(cores...),
		),
	}
}

// Debug is a zap.Logger's Debug func wrapper
func (l *Logger) Debug(format string, params ...interface{}) {
	l.log.Sugar().Debugf(format+"\n", params...)
}

// Info is a zap.Logger's Info func wrapper
func (l *Logger) Info(format string, params ...interface{}) {
	l.log.Sugar().Infof(format+"\n", params...)
}

// Warn is a zap.Logger's Warn func wrapper
func (l *Logger) Warn(format string, params ...interface{}) {
	l.log.Sugar().Warnf(format+"\n", params...)
}

// Error is a zap.Logger's Error func wrapper
func (l *Logger) Error(format string, params ...interface{}) {
	l.log.Sugar().Errorf(format+"\n", params...)
}

// Panic is a zap.Logger's Panic func wrapper
func (l *Logger) Panic(format string, params ...interface{}) {
	l.log.Sugar().Panicf(format+"\n", params...)
}

// TestLogger is Interface impl for tests
type TestLogger struct{}

// Debug for Interface impl
func (tl *TestLogger) Debug(_ string, _ ...interface{}) {}

// Info for Interface impl
func (tl *TestLogger) Info(_ string, _ ...interface{}) {}

// Warn for Interface impl
func (tl *TestLogger) Warn(_ string, _ ...interface{}) {}

// Error for Interface impl
func (tl *TestLogger) Error(_ string, _ ...interface{}) {}

// Panic for Interface impl
func (tl *TestLogger) Panic(_ string, _ ...interface{}) {}
