package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"os"
)

var standardLogger = logrus.New()
var fileLoggerEntry *logrus.Entry
var fileLogger = logrus.New()

func init() {
	configureFileLogger()
}

func configureFileLogger() {
	addFields()
	configureOutputFile()
}

func addFields() {
	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		serviceName = "goapp"
	}
	fileLoggerEntry = fileLogger.WithField("service.name", serviceName).WithField("event.dataset", serviceName)
}

func configureOutputFile() {
	loggingFileName := os.Getenv("LOGGING_FILE_NAME")

	if loggingFileName == "" {
		loggingFileName = "/tmp/logrus.log"
	}
	file, err := os.OpenFile(loggingFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		fileLogger.Out = file
		fileLogger.Formatter = &ecslogrus.Formatter{}
	} else {
		standardLogger.Info(fmt.Sprintf("Failed to log to file '%s', using default stderr", loggingFileName))
	}
}

func Info(args ...interface{}) {
	standardLogger.Info(args...)
	fileLoggerEntry.Info(args...)
}

func Error(args ...interface{}) {
	// Alternative: Distinct error loggers
	standardLogger.SetReportCaller(true)
	standardLogger.Error(args...)
	standardLogger.SetReportCaller(false)

	fileLogger.SetReportCaller(true)
	fileLoggerEntry.Error(args...)
	fileLogger.SetReportCaller(false)
}
