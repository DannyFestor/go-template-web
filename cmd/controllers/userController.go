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
		c.response.View(w, r, "user.dashboard", &templates.Data{})
	})
}
