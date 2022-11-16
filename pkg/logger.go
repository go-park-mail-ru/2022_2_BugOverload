package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func NewLogger(config *pkg.Logger) (*logrus.Logger, func() error) {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.Fatal(err)
	}

	logger := logrus.New()

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		logrus.Fatalf("err get time zone")
	}

	currentTime := time.Now().In(location)

	formatted := config.LogAddr + fmt.Sprintf("Date:%d.%02d.%02d__Time:%02d:%02d:%02d",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second()) + ".log"

	err = os.MkdirAll(config.LogAddr, 0777)
	if err != nil {
		if os.IsExist(err) {
			info, errInfo := os.Stat(config.LogAddr)
			if errInfo != nil {
				logrus.Fatalf("Error creating directory for logs: [%s]", errInfo)
			}
			if !info.IsDir() {
				logrus.Fatalf("Error: path exists but is not directory")
			}
		}
		logrus.Fatalf("Error creating directory for logs: [%s]", err)
	}

	f, err := os.OpenFile(formatted, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	logger.SetOutput(f)
	logger.Writer()
	logger.SetLevel(level)

	formatter := &logrus.TextFormatter{}
	formatter.DisableQuote = true

	logger.SetFormatter(formatter)

	return logger, f.Close
}
