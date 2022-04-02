package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
)

type singleton struct {
	Log *logrus.Logger
	sync.RWMutex
}

var instance singleton

func init() {
	_ = InitLogger("", "debug", true)
}

func GetInstance() *singleton {
	return &instance
}

func (this *singleton) Debug(message interface{}) {
	instance.Log.Debug(message)
}

func (this *singleton) Debugf(message string, args ...interface{}) {
	this.Debug(fmt.Sprintf(message, args...))
}

func (this *singleton) Info(message interface{}) {
	instance.Log.Info(message)
}

func (this *singleton) Infof(message string, args ...interface{}) {
	this.Info(fmt.Sprintf(message, args...))
}

func (this *singleton) Error(message interface{}) {
	instance.Log.Error(message)
}

func (this *singleton) Errorf(message string, args ...interface{}) {
	this.Error(fmt.Sprintf(message, args...))
}

func parseLogLevel(logLevel string) (logrus.Level, error) {

	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "debug":
		return logrus.DebugLevel, nil
	case "info":
		return logrus.InfoLevel, nil
	default:
		return logrus.InfoLevel, fmt.Errorf("Log level given %s logLevel it is invalid", logLevel)
	}
}

func InitLogger(customPath string, level string, callerTrace bool) error {
	logPath := ""
	if path := os.Getenv("METRICS_LOG"); path != "" {
		logPath = path
	} else {
		homePath, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		logPath = homePath
	}
	if customPath != "" {
		logPath = customPath
	}

	instance.Log = logrus.New()
	logLevel, err := parseLogLevel(level)
	if err != nil {
		return err
	}
	instance.Log.SetLevel(logLevel)
	instance.Log.SetReportCaller(callerTrace)

	if callerTrace {
		instance.Log.SetReportCaller(true)
	}
	file, err := os.OpenFile(logPath+"/metrics.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		instance.Log.Out = file
	} else {
		instance.Log.Info("Failed to log to file, using default stderr")
	}
	return nil
}
