package ui

import (
	"os"

	"github.com/charmbracelet/log"
)

var (
	// Logger is the global logger instance
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           log.InfoLevel,
		ReportTimestamp: false,
		ReportCaller:    false,
	})
)

// Initialize sets up the logger with the appropriate configuration
func Initialize() {
	// Configure the logger with custom settings if needed
	Logger.SetLevel(log.InfoLevel)
}

// Success logs a success message
func LogSuccess(msg string, args ...interface{}) {
	Logger.Info(IconSuccess+" "+msg, args...)
}

// Error logs an error message
func LogError(msg string, args ...interface{}) {
	Logger.Error(msg, args...)
}

// Warning logs a warning message
func LogWarning(msg string, args ...interface{}) {
	Logger.Warn(IconWarning+" "+msg, args...)
}

// Info logs an informational message
func LogInfo(msg string, args ...interface{}) {
	Logger.Info(msg, args...)
}

// Title logs a title message
func LogTitle(msg string, args ...interface{}) {
	Logger.Info(IconSparkle+" "+msg, args...)
}

// Empty logs a message for empty results
func LogEmpty(msg string, args ...interface{}) {
	Logger.Warn(msg, args...)
}

// Debug logs a debug message
func LogDebug(msg string, args ...interface{}) {
	Logger.Debug(msg, args...)
}
