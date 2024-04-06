package controllers

import (
	"net/http"
	"strconv"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/internals/pages"
)

type ErrorController struct {
	app *config.Application
}

func (e *ErrorController) Handle(statusCode int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			StatusCode int
		}
		d := data{
			StatusCode: statusCode,
		}
		err := pages.Render(e.app, w, "errors."+strconv.Itoa(statusCode), d)
		if err != nil {
			err := pages.Render(e.app, w, "errors.404", d)
			if err != nil {
				e.app.Logger.Error(err.Error())
				w.Write([]byte("Something went wrong"))
				return
			}
		}
	})
}
