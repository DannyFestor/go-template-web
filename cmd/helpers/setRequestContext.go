package helpers

import (
	"context"
	"net/http"
)

func SetRequestContext(r *http.Request, key string, value any) *http.Request {
	ctx := context.WithValue(r.Context(), key, value)
	return r.WithContext(ctx)
}
