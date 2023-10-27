package wails

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// MultiLogger is a utility to log messages to a number of destinations
type MultiLogger struct {
	filename    string
	logFile     *os.File
	multiWriter io.Writer
}

// NewMultiLogger creates a new Logger.
func NewMultiLogger(filename string) logger.Logger {
	return &MultiLogger{
		filename: filename,
	}
}

// Print works like Sprintf.
func (l *MultiLogger) Print(message string) {
	if l.multiWriter == nil {
		f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalln("OpenFile Error:", err.Error())
			log.Print(message)
			return
		}
		l.logFile = f
		l.multiWriter = io.MultiWriter(f, os.Stdout)
	}
	if _, err := l.multiWriter.Write([]byte(message)); err != nil {
		if strings.Contains(err.Error(), "write /dev/stdout: The handle is invalid") {
			l.multiWriter = l.logFile
			l.Print(message)
			return
		}
		log.Println(err.Error())
		log.Println(message)
		return
	}
}

func (l *MultiLogger) Println(message string) {
	dataTime := time.Now().Format(time.DateTime)
	l.Print(dataTime + " | " + message + "\n")
}

// Trace level logging. Works like Sprintf.
func (l *MultiLogger) Trace(message string) {
	l.Println("TRACE | " + message)
}

// Debug level logging. Works like Sprintf.
func (l *MultiLogger) Debug(message string) {
	l.Println("DEBUG | " + message)
}

// Info level logging. Works like Sprintf.
func (l *MultiLogger) Info(message string) {
	l.Println("INFO  | " + message)
}

// Warning level logging. Works like Sprintf.
func (l *MultiLogger) Warning(message string) {
	l.Println("WARN  | " + message)
}

// Error level logging. Works like Sprintf.
func (l *MultiLogger) Error(message string) {
	l.Println("ERROR | " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *MultiLogger) Fatal(message string) {
	l.Println("FATAL | " + message)
	os.Exit(1)
}
