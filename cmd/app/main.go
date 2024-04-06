package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/cmd/routes"
	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

func main() {
	conf := config.NewConfig()

	templateCache, err := templates.NewTemplateCatche(filepath.Join("resources", "views"))
	if err != nil {
		panic(err)
	}
	fmt.Println(templateCache)
	app := config.NewApplication(templateCache)

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
	err = srv.ListenAndServe()
	if err != nil {
		panic("Not working!!")
	}
}
