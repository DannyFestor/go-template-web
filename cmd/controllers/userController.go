package controllers

import (
	"bytes"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
)

type UserController struct {
	App *config.Application
}

func (c *UserController) Dashboard() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tmpl, ok := app.Templates["user/dashboard.page.tmpl"]
		tmpl, ok := c.App.Templates["user.dashboard"]
		// tmpl, ok := c.App.Templates["home"]
		if !ok {
			return
		}

		// TODO: Render Help Function
		type data struct{}
		d := data{}
		buf := new(bytes.Buffer)
		err := tmpl.Execute(buf, d)
		if err != nil {
			c.App.Logger.Error(err.Error())
			w.Write([]byte("Something went wrong"))
			return
		}

		buf.WriteTo(w)
	})
}
