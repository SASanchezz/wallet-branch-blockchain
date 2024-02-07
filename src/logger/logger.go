package logger

import (
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	Path string
}

func (l *Logger) Log(message string) {
	dir := filepath.Dir(l.Path)

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	// Open or create the log file
	logFile, err := os.OpenFile(l.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// Create a new logger with the log file as the output
	logger := log.New(logFile, "", 0)

	// Log the message
	logger.Println(message)
}
