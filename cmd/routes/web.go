package routes

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/controllers"
)

func web(c *controllers.Controllers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /{error}", c.ErrorController.Handle(404))
	mux.Handle("GET /", c.HomeController.Index())
	mux.Handle("GET /dashboard", c.UserController.Dashboard())

	return mux
}
