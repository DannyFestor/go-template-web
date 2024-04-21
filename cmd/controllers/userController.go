package controllers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/response"
	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type UserController struct {
	response *response.Response
}

func (c *UserController) Dashboard() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := c.response.View(w, r, "user.dashboard", &templates.Data{})
		if err != nil {
			// c.app.Logger.Error(err.Error()) // TODO: find a good way to log an error
			w.Write([]byte("Something went wrong"))
		}
	})
}
