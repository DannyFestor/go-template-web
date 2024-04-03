package main

import (
	"fmt"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/routes"
)

func main() {
	handler := routes.Get()

	srv := http.Server{
		Addr:    ":4000",
		Handler: handler,
	}

	fmt.Println("Running on Port :4000")
	err := srv.ListenAndServe()
	if err != nil {
		panic("Not working!!")
	}
}
