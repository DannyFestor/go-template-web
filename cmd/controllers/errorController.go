package controllers

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
)

type ErrorController struct {
	App *config.Application
}

func (e *ErrorController) Handle(statusCode int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle Error Template Not Found
		tmpl, ok := e.App.Templates["errors."+strconv.Itoa(statusCode)]
		if !ok {
			tmpl, ok = e.App.Templates["errors.404"]
			if !ok {
				return
			}
		}

		type data struct {
			StatusCode int
		}
		d := data{
			StatusCode: statusCode,
		}
		buf := new(bytes.Buffer)
		err := tmpl.Execute(buf, d)
		if err != nil {
			e.App.Logger.Error(err.Error())
			w.Write([]byte("Something went wrong"))
			return
		}

		buf.WriteTo(w)
	})
}
