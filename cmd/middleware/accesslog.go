package middleware

import (
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (mw *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(ww, r)

		isHtmxRequest := false
		if r.Header.Get("Hx-Request") == "true" {
			isHtmxRequest = true
		}

		mw.App.Logger.Info(r.RemoteAddr, "method", r.Method, "path", r.URL.Path, "htmx", isHtmxRequest, "time", time.Since(start))
	})
}
