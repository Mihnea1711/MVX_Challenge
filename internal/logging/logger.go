package logging

import (
	logger "github.com/multiversx/mx-chain-logger-go"
)

var log logger.Logger

// SetupLogger creates and configures a new logger instance.
func SetupLogger() logger.Logger {
	log = logger.GetOrCreate("my_logger")
	log.SetLevel(logger.LogTrace)
	return log
}

// GetLogger returns the logger instance.
func GetLogger() logger.Logger {
	return log
}

// Initialize the logger during package initialization
func init() {
	SetupLogger()
}
