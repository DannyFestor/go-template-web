package routes

import (
	"fmt"
	"net/http"
)

func api() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		fmt.Println("Query!")
		fmt.Println(query)

		if query.Get("name") == "" {
			w.Write([]byte("Hello World from API"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Hello %s, from API", query.Get("name"))))
	}))

	return mux
}

