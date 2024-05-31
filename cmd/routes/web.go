package routes

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/DannyFestor/go-template-web.git/cmd/controllers"
	"github.com/DannyFestor/go-template-web.git/cmd/middleware"
	"github.com/DannyFestor/go-template-web.git/resources"
)

func web(mw *middleware.Middleware, c *controllers.Controllers) (http.Handler, error) {
	mux := http.NewServeMux()
	webMiddlewares := middleware.Chain(
		mw.HandleHtmxRequest,
	)

	mux.Handle("GET /dashboard", c.UserController.Dashboard())

	mux.Handle("GET /", handleDefault(c.HomeController.Index(), c.ErrorController.Handle(404)))

	return webMiddlewares(mux), nil
}

func handleDefault(c http.Handler, e http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			public, err := fs.Sub(resources.EmbeddedFiles, "public")
			if err != nil {
				e.ServeHTTP(w, r)
				return
			}

			file, err := public.Open(strings.TrimPrefix(r.URL.Path, "/"))
			if err != nil {
				e.ServeHTTP(w, r)
				return
			}
			defer file.Close()

			http.FileServer(http.FS(public)).ServeHTTP(w, r)
			return
		}

		c.ServeHTTP(w, r)
	})
}
