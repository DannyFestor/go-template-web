package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		fmt.Println("Query!")
		fmt.Println(query)

		if query.Get("name") == "" {
			w.Write([]byte("Hello World"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Hello %s", query.Get("name"))))
	}))

	srv := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	fmt.Println("Running on Port :4000")
	err := srv.ListenAndServe()
	if err != nil {
		panic("Not working!!")
	}
}
