package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/cmd/routes"
)

func main() {
	conf := config.NewConfig()
	app := config.NewApplication()

	handler := routes.Get(app)

	srv := http.Server{
		Addr:         conf.Port,
		Handler:      handler,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelWarn),
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	app.Logger.Info(fmt.Sprintf("Running on Port %s", conf.Port))
	err := srv.ListenAndServe()
	if err != nil {
		panic("Not working!!")
	}
}
