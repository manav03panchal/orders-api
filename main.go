package main

import "fmt"
import "net/http"

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	las := server.ListenAndServe()
	if las != nil {
		fmt.Println("failed to listen to server", las)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
