package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/cmd/routes"
)

// TODO: Add a run function?
func main() {
	conf := config.NewConfig()

	app, err := config.NewApplication(conf)
	if err != nil {
		app.Logger.Error("Main NewApplication failed", "Error", err.Error())
		os.Exit(1)
	}
	defer app.Db.Close(context.Background())

	handler, err := routes.Get(app)
	if err != nil {
		app.Logger.Error("Main routes.Get failed", "Error", err.Error())
		os.Exit(1)
	}

	srv := http.Server{
		Addr:         conf.Port,
		Handler:      handler,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelWarn),
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	app.Logger.Info(fmt.Sprintf("Running on Port %s", conf.Port))
	err = srv.ListenAndServe()
	if err != nil {
		app.Logger.Error("Main src.ListenAndServe failed", "Error", err.Error())
		os.Exit(1)
	}
}
