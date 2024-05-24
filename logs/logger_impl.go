package logs

import (
	"log"
	"os"
	"sync"
)

// SimpleLogger is a basic implementation of the Logger interface
type SimpleLogger struct {
	logger *log.Logger
}

var (
	instance *SimpleLogger
	once     sync.Once
)

// GetLogger returns a singleton instance of SimpleLogger
func NewSimpleLogger() *SimpleLogger {
	once.Do(func() {
		file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		instance = &SimpleLogger{
			logger: log.New(file, "", log.LstdFlags),
		}
	})
	return instance
}

func (l *SimpleLogger) Info(msg string, args ...interface{}) {
	l.logger.Printf("[INFO] "+msg, args...)
}

func (l *SimpleLogger) Error(msg string, args ...interface{}) {
	l.logger.Printf("[ERROR] "+msg, args...)
}
