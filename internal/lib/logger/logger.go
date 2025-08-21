package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
	})

	Log.SetLevel(logrus.InfoLevel)
	Log.SetOutput(os.Stdout)
}

func WithError(err error) *logrus.Entry {
	return Log.WithError(err)
}
