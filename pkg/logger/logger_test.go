package logger_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/fsvxavier/unico/pkg/logger"
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

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouterLogger() *gin.Engine {
	var logg logger.GenericLogger
	logg.Module = "teste"
	logg.GetLogger()

	router := gin.New()
	router.Use(logg.Logger(logg.Log.Logger))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(http.StatusOK, "patch")
	})
	router.OPTIONS("/", func(c *gin.Context) {
		c.String(http.StatusOK, "options")
	})
	return router
}
