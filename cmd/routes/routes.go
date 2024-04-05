package routes

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/cmd/middleware"
)

func Get(app *config.Application) http.Handler {
	mw := middleware.Middleware{
		App: app,
	}

	middlewares := middleware.Chain(
		mw.Log,
	)

	mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", api()))
	mux.Handle("/", web())

	return middlewares(mux)
}
