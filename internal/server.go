package internal

import (
	"io"
	"net/http"
	"strconv"
)

var shortToLongMapping = make(map[string]string, 0)

func handleRoot(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(request.Body)
	shortUrl := doShort(string(body))
	writer.WriteHeader(201)
	writer.Write([]byte(shortUrl))
}

func handleId(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Location", shortToLongMapping[request.PathValue("id")])
	writer.WriteHeader(307)
}

func doShort(url string) string {
	short := strconv.Itoa(len(url) * 123)
	shortToLongMapping[short] = url
	return short
}

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/{id}", handleId)

	http.ListenAndServe("localhost:8080", mux)
}
