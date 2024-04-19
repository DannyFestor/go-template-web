package routes

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/DannyFestor/go-template-web.git/cmd/controllers"
	"github.com/DannyFestor/go-template-web.git/resources"
)

func web(c *controllers.Controllers) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	// mux.Handle("GET /", assets)
	// mux.Handle("GET /", c.HomeController.Index())
	mux.Handle("GET /dashboard", c.UserController.Dashboard())

	mux.Handle("GET /", handleDefault(c))

	// mux.Handle("GET /*", c.ErrorController.Handle(404))

	return mux, nil
}

func handleDefault(c *controllers.Controllers) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			public, err := fs.Sub(resources.EmbeddedFiles, "public")
			if err != nil {
				c.ErrorController.Handle(404).ServeHTTP(w, r)
				return
			}

			file, err := public.Open(strings.TrimPrefix(r.URL.Path, "/"))
			if err != nil {
				fmt.Println(err)
				c.ErrorController.Handle(404).ServeHTTP(w, r)
				return
			}
			defer file.Close()

			http.FileServer(http.FS(public)).ServeHTTP(w, r)
			return
		}

		c.HomeController.Index().ServeHTTP(w, r)
	})
}
