package internal

import "fmt"

type URLStorage struct {
	mapping map[string]string
}

func (storage URLStorage) Add(shortURL, fullURL string) {
	storage.mapping[shortURL] = fullURL
}

func (storage URLStorage) Get(shortURL string) (string, error) {
	urlID, success := storage.mapping[shortURL]
	if !success {
		return "", fmt.Errorf("url %s dont have full version", shortURL)
	}
	return urlID, nil
}

func NewURLStorage() *URLStorage {
	return &URLStorage{mapping: make(map[string]string, 0)}
}
