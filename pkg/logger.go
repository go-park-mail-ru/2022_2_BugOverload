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

	currentTime := time.Now().In(time.UTC)

	formatted := config.LogAddr + "__" + fmt.Sprintf("Date:%d-.%02d.%02d__Time:%02d:%02d:%02d",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second()) + ".log"

	f, err := os.OpenFile(formatted, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	logger.SetOutput(f)
	logger.Writer()
	logger.SetLevel(level)
	// logger.SetFormatter(&logrus.JSONFormatter{})

	return logger, f.Close
}
