package internal

import (
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()
	storage := URLStorage{mapping: make(map[string]string, 0)}
	mux.HandleFunc("/", RootHandler(storage))
	mux.HandleFunc("/{id}", IDHandler(storage))

	http.ListenAndServe(":8080", mux)
}
