package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	configM "github.com/prayogatriady/ecommerce-module/config"
	"github.com/prayogatriady/ecommerce-order/utils/constant"
	"github.com/sirupsen/logrus"
)

var (
	appName  = configM.String("app.name", "")
	tz       = configM.String("app.timezone", "")
	logDir   = configM.String("logger.dir", "")
	logLevel = configM.String("logger.level", "")

	Logger *logrus.Logger
)

func InitLogger() {

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	loc, _ := time.LoadLocation(tz)

	go dailyLog(loc)

	InitLogrus(loc)

}

func dailyLog(loc *time.Location) {

	nextDay := time.Now().In(loc).Add(24 * time.Hour)
	nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, loc)
	sleepDuration := time.Until(nextDay)
	ticker := time.NewTicker(sleepDuration)

	quit := make(chan bool)

	for {
		select {
		case <-ticker.C:
			InitLogrus(loc)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func InitLogrus(loc *time.Location) {

	currentDate := time.Now().In(loc).Format(constant.YYYYMMDD)

	logFilename := fmt.Sprintf("%s_%s.log", appName, currentDate)
	logFullDir := filepath.Join(logDir, logFilename)
	logFile, err := os.OpenFile(logFullDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
	}

	Logger = logrus.New()
	Logger.SetOutput(logFile)

	levelMap := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	Logger.SetLevel(levelMap[logLevel])
}
