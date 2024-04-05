package middleware

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
)

type Middleware struct {
	App *config.Application
}

type MWStack func(http.Handler) http.Handler

func Chain(middlewares ...MWStack) MWStack {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			mw := middlewares[i]
			next = mw(next)
		}

		return next
	}
}
