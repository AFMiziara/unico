package logger

import (
	"os"
	"time"

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

	if os.Getenv("LOG_FILE_PATH") != "" {
		currentTime := time.Now()
		f, err := os.OpenFile(os.Getenv("LOG_FILE_PATH")+"/logrus_"+currentTime.Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			rootLogger.SetOutput(f)
		}
	}

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
