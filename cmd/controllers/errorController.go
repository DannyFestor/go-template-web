package controllers

import (
	"net/http"
	"strconv"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type ErrorController struct {
	app *config.Application
}

func (e *ErrorController) Handle(statusCode int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := &templates.Data{
			StatusCode: statusCode,
		}
		err := e.app.Response.View(w, r, "errors."+strconv.Itoa(statusCode), d)
		if err != nil {
			err := e.app.Response.View(w, r, "errors.404", d)
			if err != nil {
				e.app.Logger.Error(err.Error())
				w.Write([]byte("Something went wrong"))
				return
			}
		}
	})
}
