package wails

import (
	"github.com/wailsapp/wails/v2/pkg/logger"
	"log"
	"os"
	"time"
)

// FileLogger is a utility to log messages to a number of destinations
type FileLogger struct {
	filename string
}

// NewFileLogger creates a new Logger.
func NewFileLogger(filename string) logger.Logger {
	return &FileLogger{
		filename: filename,
	}
}

// Print works like Sprintf.
func (l *FileLogger) Print(message string) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(message); err != nil {
		f.Close()
		log.Fatal(err)
	}
	f.Close()
}

func (l *FileLogger) Println(message string) {
	dataTime := time.Now().Format(time.DateTime)
	l.Print(dataTime + " | " + message + "\n")
}

// Trace level logging. Works like Sprintf.
func (l *FileLogger) Trace(message string) {
	l.Println("TRACE | " + message)
}

// Debug level logging. Works like Sprintf.
func (l *FileLogger) Debug(message string) {
	l.Println("DEBUG | " + message)
}

// Info level logging. Works like Sprintf.
func (l *FileLogger) Info(message string) {
	l.Println("INFO  | " + message)
}

// Warning level logging. Works like Sprintf.
func (l *FileLogger) Warning(message string) {
	l.Println("WARN  | " + message)
}

// Error level logging. Works like Sprintf.
func (l *FileLogger) Error(message string) {
	l.Println("ERROR | " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *FileLogger) Fatal(message string) {
	l.Println("FATAL | " + message)
	os.Exit(1)
}
