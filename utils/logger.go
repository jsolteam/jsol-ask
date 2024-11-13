package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// LogType represents the type of log (e.g., INFO, WARN, ERROR, PANIC)
type LogType string

const (
	InfoLog  LogType = "INFO"
	WarnLog  LogType = "WARN"
	ErrorLog LogType = "ERROR"
	PanicLog LogType = "PANIC"
)

// ModuleTag represents the module generating the log (e.g., SYSTEM, DATABASE, HANDLER, SERVICE)
type ModuleTag string

const (
	SystemModule   ModuleTag = "SYSTEM"
	DatabaseModule ModuleTag = "DATABASE"
	HandlerModule  ModuleTag = "HANDLER"
	ServiceModule  ModuleTag = "SERVICE"
)

// Logger handles logging to separate files and optionally to a database
type Logger struct {
	//db *sql.DB
	mu sync.Mutex
}

var (
	instance *Logger
	once     sync.Once
)

// InitLogger initializes the singleton Logger with an optional database connection.
// It should be called only once, typically at the start of the application.
func InitLogger() {
	once.Do(func() {
		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			log.Panicf("Failed to create logs directory: %v", err)
		}
		instance = &Logger{}
	})
}

// Log is a global function for logging messages using the singleton Logger instance.
// It logs the message to both the appropriate log file and the database with the given log type and module tag.
func Log(logType LogType, module ModuleTag, message string) {
	if instance == nil {
		log.Panic("Logger is not initialized. Call InitLogger() before using Log().")
	}
	instance.log(logType, module, message)
}

// log is the actual logging method used by the global Log function
func (l *Logger) log(logType LogType, module ModuleTag, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := formatLogEntry(timestamp, logType, module, message)

	// Write log to file for the specific log type
	if err := l.writeLogToFile(logType, logEntry); err != nil {
		log.Printf("Failed to write log to file: %v", err)
	}

	// Optionally write log to database if db is set
	//if l.db != nil {
	//	if err := l.saveLogToDatabase(logType, module, message, timestamp); err != nil {
	//		log.Printf("Failed to save log to database: %v", err)
	//	}
	//}
}

// writeLogToFile writes a log entry to a specific file based on log type.
func (l *Logger) writeLogToFile(logType LogType, entry string) error {
	logFilePath := filepath.Join("logs", fmt.Sprintf("%s.log", logType))
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file %s: %v", logFilePath, err)
	}
	defer file.Close()

	if _, err := file.WriteString(entry + "\n"); err != nil {
		return fmt.Errorf("failed to write to log file %s: %v", logFilePath, err)
	}

	return nil
}

// formatLogEntry formats the log entry string with timestamp, log type, module, and message.
func formatLogEntry(timestamp string, logType LogType, module ModuleTag, message string) string {
	return fmt.Sprintf("%s [%s] [%s] %s", timestamp, logType, module, message)
}

// saveLogToDatabase inserts the log entry into the database.
//func (l *Logger) saveLogToDatabase(logType LogType, module ModuleTag, message, timestamp string) error {
//	_, err := l.db.Exec(`
//		INSERT INTO logs (timestamp, log_type, module, message)
//		VALUES (?, ?, ?, ?)`,
//		timestamp, logType, module, message)
//	return err
//}

// Close closes the database connection (if any).
//func (l *Logger) Close() error {
//	if l.db != nil {
//		return l.db.Close()
//	}
//	return nil
//}
