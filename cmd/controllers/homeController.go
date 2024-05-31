package controllers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
	"github.com/DannyFestor/go-template-web.git/cmd/helpers"
	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type HomeController struct {
	app *config.Application
}

func (c *HomeController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Hx-Request") == "true" {
			r = helpers.SetRenderBlock(r, "test")
		}
		c.app.Response.View(w, r, "home", &templates.Data{})
	})
}
