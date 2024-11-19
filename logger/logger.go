package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	logTimeLayout = "02-01-2006_03:04:05"

	dirPermissions  = 0755
	filePermissions = 0600
	outDir          = "out/log"
)

// Logger errors
var (
	ErrOpen   = errors.New("couldn't open log file")
	ErrCreate = errors.New("couldn't create log directory")
)

// InitLogger initiate auth service logger
func InitLogger(serviceName string) *log.Logger {
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outDir, dirPermissions); err != nil {
			log.Fatalf("%v - %v: %v", ErrCreate, outDir, err)
		}
	}

	logFileName := fmt.Sprintf("%s/%s_%s.log", outDir, serviceName, time.Now().Format(logTimeLayout))

	out, err := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR, filePermissions) //nolint: gosec
	if err != nil {
		log.Fatalf("%v - %v: %v", ErrOpen, outDir, err)
	}

	return log.New(out, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile)
}
