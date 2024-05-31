package middleware

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/helpers"
)

func (mw *Middleware) HandleHtmxRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var renderedBlock string
		if r.Header.Get("Hx-Request") == "true" {
			renderedBlock = "body"
		} else {
			renderedBlock = "base"
		}

		r = helpers.SetRenderBlock(r, renderedBlock)

		next.ServeHTTP(w, r)
	})
}
