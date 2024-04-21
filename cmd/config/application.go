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

	ContextKeys *ContextKeys
}

func NewApplication() (*Application, error) {
	loggerOptions := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	}
	logger := slog.New(tint.NewHandler(os.Stdout, loggerOptions))

	contextKeys := &ContextKeys{
		Htmx: "HtmxRequest",
	}

	response, err := response.NewResponse(string(contextKeys.Htmx))
	if err != nil {
		return nil, err
	}

	app := &Application{
		Logger: logger,

		Response: response,

		ContextKeys: contextKeys,
	}

	return app, nil
}
