package internal

import (
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()
	storage := NewURLStorage()
	mux.HandleFunc("/", RootHandler(storage))
	mux.HandleFunc("/{id}", IDHandler(storage))

	http.ListenAndServe(":8080", mux)
}
