package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/prayogatriady/ecommerce-order/utils/config"
	"github.com/prayogatriady/ecommerce-order/utils/constant"
	"github.com/sirupsen/logrus"
)

var (
	env    = config.Env
	Logger *logrus.Logger
)

func InitLogger() {

	if _, err := os.Stat(env.Logger.Dir); os.IsNotExist(err) {
		os.Mkdir(env.Logger.Dir, os.ModePerm)
	}

	loc, _ := time.LoadLocation(env.App.Timezone)

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

	logFilename := fmt.Sprintf("%s_%s.log", env.App.Name, currentDate)
	logFullDir := filepath.Join(env.Logger.Dir, logFilename)
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

	Logger.SetLevel(levelMap[env.Logger.Level])
}
