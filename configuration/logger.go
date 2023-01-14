package configuration

import (
	"golang-todo-app/exception"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		exception.PanicLogging(err)
	}

	date := time.Now()
	file, err := os.OpenFile("logs/log_"+date.Format("01-02-2006")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicLogging(err)
	if err == nil {
		logger.SetOutput(file)
	}
	return logger
}

func NewLoggerConfig() logger.Config {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		exception.PanicLogging(err)
	}

	date := time.Now()
	file, err := os.OpenFile("logs/log_"+date.Format("01-02-2006")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicLogging(err)

	return logger.Config{
		Output: file,
	}
}

/*
ref:
- https://dev.to/koddr/go-fiber-by-examples-working-with-middlewares-and-boilerplates-3p0m#explore-logging-middleware
*/
