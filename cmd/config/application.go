package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/DannyFestor/go-template-web.git/internals/response"
	"github.com/lmittmann/tint"
)

// holds application information
type Application struct {
	// config
	Logger *slog.Logger
	// db
	// mail
	// session
	Response *response.Response
}

func NewApplication() (*Application, error) {
	loggerOptions := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	}
	logger := slog.New(tint.NewHandler(os.Stdout, loggerOptions))

	response, err := response.NewResponse()
	if err != nil {
		return nil, err
	}

	app := &Application{
		Logger: logger,

		Response: response,
	}

	return app, nil
}
