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

	ILog *logrus.Logger
	ELog *logrus.Logger
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

	logFilenameInfo := fmt.Sprintf("%s_%s_INFO.log", currentDate, appName)
	logFilenameError := fmt.Sprintf("%s_%s_ERROR.log", currentDate, appName)

	logFullDirInfo := filepath.Join(logDir, logFilenameInfo)
	logFullDirError := filepath.Join(logDir, logFilenameError)

	logFileInfo, err := os.OpenFile(logFullDirInfo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
	}
	logFileError, err := os.OpenFile(logFullDirError, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
	}

	levelMap := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	ILog = logrus.New()
	ILog.SetOutput(logFileInfo)
	ILog.SetLevel(levelMap[logLevel])

	ELog = logrus.New()
	ELog.SetOutput(logFileError)
	ELog.SetLevel(levelMap[logLevel])
}
