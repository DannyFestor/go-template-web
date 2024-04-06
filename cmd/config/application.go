package config

import (
	"html/template"
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

	Templates map[string]*template.Template
}

func NewApplication(templateCache map[string]*template.Template) *Application {
	loggerOptions := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	}
	logger := slog.New(tint.NewHandler(os.Stdout, loggerOptions))

	return &Application{
		Logger:    logger,
		Templates: templateCache,
	}
}
