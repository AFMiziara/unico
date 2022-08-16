package logger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Override for testing
var osHostname = os.Hostname

var (
	rootLogger *logrus.Logger
)

// GenericLogger represents log struct
type GenericLogger struct {
	Log           *logrus.Entry
	Hostname      string
	Module        string
	OperationName string
}

func initLogger() *logrus.Logger {
	rootLogger = logrus.New()
	rootLogger.SetNoLock()
	rootLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	rootLogger.SetLevel(getLogLevel("LOG_LEVEL"))
	return rootLogger
}

func (g *GenericLogger) GetHostname() (string, error) {
	host, err := osHostname()
	if err != nil {
		return "unknown", err
	}
	return host, nil
}

// GetLogger return a initialized log
func (g *GenericLogger) GetLogger() {
	if rootLogger == nil {
		initLogger()
	}

	hostname := "unknown"
	hostname, _ = g.GetHostname()
	g.Hostname = hostname

	//g.Hostname = Hostname
	g.Log = rootLogger.WithFields(logrus.Fields{
		"environment": os.Getenv("ENV"),
		"hostname":    hostname,
		"version":     os.Getenv("VERSION"),
		"app":         os.Getenv("APP"),
		"squad":       os.Getenv("SQUAD"),
		"tribe":       os.Getenv("TRIBE"),
		"module":      g.Module,
	})
}

// LogIt log a new message to stdout
func (g *GenericLogger) LogIt(severity, message string, fields map[string]interface{}) {
	logg := g.Log
	logg = logg.WithFields(logrus.Fields{
		"severity": severity,
		"operation": logrus.Fields{
			"name": g.OperationName,
		},
	})
	if fields != nil {
		logg = logg.WithFields(fields)
	}
	switch severity {
	case "DEBUG":
		logg.Warn(message)
	case "INFO":
		logg.Info(message)
	case "WARN":
		logg.Warn(message)
	case "ERROR":
		logg.Error(message)
	case "FATAL":
		logg.Fatal(message)
	case "PANIC":
		logg.Panic(message)
	default:
		logg.Info(message)
	}
}

func getLogLevel(envVariable string) logrus.Level {
	switch os.Getenv(envVariable) {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	case "PANIC":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

var timeFormat = "02/Jan/2006:15:04:05 -0700"

// Logger is the logrus logger handler
func (g *GenericLogger) Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := map[string]interface{}{
			"http": map[string]interface{}{
				"latency":        latency, // time to process
				"request_method": c.Request.Method,
				"status_code":    statusCode,
				"path":           path,
			},
			"clientIP":   clientIP,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		}

		logger := GenericLogger{}
		logger.Module = "server"
		logger.GetLogger()

		if len(c.Errors) > 0 {
			logger.LogIt("ERROR", c.Errors.ByType(gin.ErrorTypePrivate).String(), entry)
		} else {

			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			msg := fmt.Sprintf("%s - %s [%s] \"%s %d\" %d %s \"%s\" (%dms) - body-send: %v", clientIP, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency, string(bodyBytes))

			switch code := statusCode; {
			case code > 499:
				logger.LogIt("ERROR", msg, entry)
			case code > 399:
				logger.LogIt("WARN", msg, entry)
			default:
				if path != "/health" {
					logger.LogIt("INFO", msg, entry)
				}
			}
		}
	}
}
