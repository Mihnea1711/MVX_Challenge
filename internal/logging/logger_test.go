package logging

import (
	"testing"

	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/stretchr/testify/require"
)

func TestSetupLogger(t *testing.T) {
	log := SetupLogger()
	require.NotNil(t, log)

	log.Trace("a trace message")
	log.Debug("a debug message")
	log.Info("an information message")
	log.Warn("a warning message")
	log.Error("an error message")

	log.Log(logger.LogTrace, "a second trace message")
	log.Log(logger.LogDebug, "a second debug message")
	log.Log(logger.LogInfo, "a second information message")
	log.Log(logger.LogWarning, "a second warning message")
	log.Log(logger.LogError, "a second error message")
	log.Log(logger.LogNone, "this message should not be printed")
}

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	require.NotNil(t, log)

	log.Info("testing GetLogger function")
}
