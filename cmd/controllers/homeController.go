package controllers

import (
	"fmt"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/response"
	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type HomeController struct {
	response *response.Response
}

func (c *HomeController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := c.response.View(w, r, "home", &templates.Data{})
		if err != nil {
			panic(fmt.Errorf("HomeController:\n%w", err))
		}
	})
}
