package config

import (
	"log/slog"

	"github.com/DannyFestor/go-template-web.git/internals/log"
	"github.com/DannyFestor/go-template-web.git/internals/response"
	"github.com/jackc/pgx/v5"
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

	app.Logger = log.MakeLogger()
	app.Logger.Info("Logger initialized")

	response, err := response.NewResponse(app.Logger)
	if err != nil {
		return nil, err
	}
	app.Response = response
	app.Logger.Info("Response initialized")

	db, err := psql(config)
	if err != nil {
		return nil, err
	}
	app.Db = db
	app.Logger.Info("Database connection successful")

	return app, nil
}
