package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// holds application information
type Application struct {
	// config
	Logger *slog.Logger
	// db
	// mail
	// session
}

func NewApplication() *Application {
	loggerOptions := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	}
	logger := slog.New(tint.NewHandler(os.Stdout, loggerOptions))

	return &Application{
		Logger: logger,
	}
}
