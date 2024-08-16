package log

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func MakeLogger() *slog.Logger {
	loggerOptions := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	}
	logger := slog.New(tint.NewHandler(os.Stdout, loggerOptions))
	return logger
}
