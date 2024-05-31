package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/DannyFestor/go-template-web.git/internals/response"
	"github.com/jackc/pgx/v5"
	"github.com/lmittmann/tint"
)

// holds application information
type Application struct {
	// config
	Logger *slog.Logger
	Db     *pgx.Conn
	// mail
	// session
	Response *response.Response
}

func NewApplication(config *Config) (*Application, error) {
	app := &Application{}

	loggerOptions := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	}
	logger := slog.New(tint.NewHandler(os.Stdout, loggerOptions))
	app.Logger = logger
	logger.Info("Logger initialized")

	response, err := response.NewResponse(logger)
	if err != nil {
		return nil, err
	}
	app.Response = response
	logger.Info("Response initialized")

	db, err := psql(config)
	if err != nil {
		return nil, err
	}
	app.Db = db
	logger.Info("Database connection successful")

	return app, nil
}
