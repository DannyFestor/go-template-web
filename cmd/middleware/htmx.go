package middleware

import (
	"context"
	"net/http"
)

func (mw *Middleware) IsHtmxRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		htmxRequest := r.Header.Get("Hx-Request") == "true"
		ctx := context.WithValue(r.Context(), mw.App.Response.HtmxKey, htmxRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
