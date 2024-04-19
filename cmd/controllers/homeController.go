package controllers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/response"
)

type HomeController struct {
	response *response.Response
}

func (c *HomeController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type data struct{}
		err := c.response.View(w, "home", data{})
		if err != nil {
			// c.response.Error(err.Error())
			w.Write([]byte("Something went wrong"))
		}
	})
}
