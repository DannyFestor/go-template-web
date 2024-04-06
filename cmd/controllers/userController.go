package controllers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/internals/pages"
)

type UserController struct {
	app *config.Application
}

func (c *UserController) Dashboard() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type data struct{}
		err := pages.Render(c.app, w, "user.dashboard", data{})
		if err != nil {
			c.app.Logger.Error(err.Error())
			w.Write([]byte("Something went wrong"))
		}
	})
}
