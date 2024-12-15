package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	fmt.Println("running server on localhost:3000 ...")
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	las := server.ListenAndServe()
	if las != nil {
		fmt.Println("failed to listen to server", las)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
