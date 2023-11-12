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

type StandardLogger struct {
	*logrus.Logger
}

var (
	// appName  = config.Env.App.Name
	// tz       = config.Env.App.Timezone
	// logDir   = config.Env.Logger.Dir
	// logLevel = config.Env.Logger.Level

	Slog *StandardLogger
)

func InitLogger() {

	if _, err := os.Stat(configM.String("logger.dir", "")); os.IsNotExist(err) {
		os.Mkdir(configM.String("logger.dir", ""), os.ModePerm)
	}

	loc, _ := time.LoadLocation(configM.String("app.timezone", ""))

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
	logFilename := fmt.Sprintf("%s_%s.log", currentDate, configM.String("app.name", ""))
	logFullDir := filepath.Join(configM.String("logger.dir", ""), logFilename)
	logFile, err := os.OpenFile(logFullDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	Slog = &StandardLogger{
		Logger: logrus.New(),
	}

	Slog.Logger.SetOutput(logFile)
	Slog.Logger.SetLevel(levelMap[configM.String("logger.level", "")])
	Slog.Logger.SetFormatter(&logrus.JSONFormatter{})
	Slog.Logger.SetReportCaller(true)
}
