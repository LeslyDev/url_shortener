package internal

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func RootHandler(storage *URLStorage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(writer, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
			return
		}
		body, _ := io.ReadAll(request.Body)
		fullURL := string(body)
		shortURL := doShort(fullURL)
		storage.Add(shortURL, fullURL)
		writer.WriteHeader(201)
		resultURL := url.URL{Scheme: "http", Host: request.Host, Path: shortURL}
		writer.Write([]byte(resultURL.String()))
	}
}

func IDHandler(storage *URLStorage) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		urlID, err := storage.Get(request.PathValue("id"))
		if err != nil {
			http.Error(writer, "url not found", http.StatusNotFound)
			return
		}
		writer.Header().Set("Location", urlID)
		writer.WriteHeader(307)
	}
}

func doShort(url string) string {
	return strconv.Itoa(len(url) * 123)
}
