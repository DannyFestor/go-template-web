package routes

import (
	"net/http"
)

func Get() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", api()))
	mux.Handle("/", web())

	return mux
}
