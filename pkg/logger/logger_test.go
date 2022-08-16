package logger_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fsvxavier/unico/pkg/logger"
	"github.com/stretchr/testify/assert"
)

// Override for testing
var osHostname = os.Hostname

func TestLogger(t *testing.T) {
	t.Run("success-severity", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "DEBUG")

		logg := new(logger.GenericLogger)
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("DEBUG", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)
		logg.LogIt("INFO", fmt.Sprintf("[test] - Info: %s", "teste"), nil)
		logg.LogIt("WARN", fmt.Sprintf("[test] - Warn: %s", "teste"), nil)
		logg.LogIt("ERROR", fmt.Sprintf("[test] - Error: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})
	t.Run("success-loglevel-debug", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "INFO")

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("DEBUG", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})
	t.Run("success-loglevel-info", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "WARN")

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("INFO", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})
	t.Run("success-loglevel-warn", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "ERROR")

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("WARN", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})

	t.Run("success-loglevel-default", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "")

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})

	t.Run("success-logrusloglevel-fatal", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "FATAL")

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})

	t.Run("success-logrusloglevel-panic", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "PANIC")

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logg)
	})

	t.Run("success-field", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "")

		field := map[string]interface{}{"teste": ""}

		var logg logger.GenericLogger
		logg.Module = "worker"
		logg.OperationName = "Initialize"
		logg.GetLogger()

		logg.LogIt("WARN", fmt.Sprintf("[test] - Debug: %s", "teste"), field)

		assert.NotNil(t, logg)
	})
}
