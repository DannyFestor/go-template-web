package controllers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type UserController struct {
	app *config.Application
}

func (c *UserController) Dashboard() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.app.Response.View(w, r, "user.dashboard", &templates.Data{})
	})
}
