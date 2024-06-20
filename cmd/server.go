package cmd

import (
	"fmt"
	"net/http"
)

func handleRoot(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	body := request.Body
	fmt.Println(body)
	writer.WriteHeader(201)
	writer.Write([]byte("Hello"))
}

func handleId(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello"))
}

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/{id}", handleId)

	http.ListenAndServe("localhost:8080", mux)
}
