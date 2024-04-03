package routes

import (
	"net/http"
)

func Get() *http.ServeMux {
	mux := http.NewServeMux()

	apiRoutes := api()
	webRoutes := web()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRoutes))
	mux.Handle("/", webRoutes)

	return mux
}
