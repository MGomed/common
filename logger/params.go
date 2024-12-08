package logger

import (
	"strings"

	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

var params = loggerParams{
	consoleConfig: zap.NewDevelopmentEncoderConfig(),
	fileConfig:    zap.NewDevelopmentEncoderConfig(),
}

type loggerParams struct {
	logLevel zap.AtomicLevel

	consoleConfig    zapcore.EncoderConfig
	isConsoleEncoder bool

	fileConfig    zapcore.EncoderConfig
	isFileEncoder bool
}

// LogOutType is log's out type (console or file)
type LogOutType string

// LogOutType consts
const (
	ConsoleOut LogOutType = "console"
	FileOut    LogOutType = "file"
)

// SetOutput sets outs for log
func SetOutput(outs ...string) {
	for _, out := range outs {
		params.isConsoleEncoder = strings.ToLower(out) == string(ConsoleOut)
		params.isFileEncoder = out == string(FileOut)
	}
}

// LogLevel is zap's AtomicLevel wrapper
type LogLevel string

// LogLevel consts
const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Panic LogLevel = "panic"
)

func (lvl LogLevel) convertToZap() zap.AtomicLevel {
	switch lvl {
	case Debug:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case Info:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case Warn:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case Error:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case Panic:
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

// SetLogLevel sets log level for log
func SetLogLevel(level string) {
	params.logLevel = LogLevel(strings.ToLower(level)).convertToZap()
}

// TimeKey is zap's TimeKey wrapper
type TimeKey string

// TimeKey consts
const (
	TimeStamp TimeKey = "timestamp"
	TS        TimeKey = "ts"
	Time      TimeKey = "time"
)

// SetTimeKey sets time key for log
func SetTimeKey(key TimeKey) {
	params.consoleConfig.TimeKey = string(key)
	params.fileConfig.TimeKey = string(key)
}

// SetStackTraceKey sets stask trace key for log
func SetStackTraceKey(key string) {
	params.consoleConfig.StacktraceKey = key
	params.fileConfig.StacktraceKey = key
}
