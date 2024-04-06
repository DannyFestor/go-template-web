package controllers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/internals/pages"
)

type HomeController struct {
	app *config.Application
}

func (c *HomeController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type data struct{}
		err := pages.Render(c.app, w, "home", data{})
		if err != nil {
			c.app.Logger.Error(err.Error())
			w.Write([]byte("Something went wrong"))
		}
	})
}
