package middleware

import (
	"fmt"
	"net/http"
)

func (mw *Middleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				mw.App.Logger.Error(fmt.Sprintf("Recover Panic: %s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
