package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIDHandler(t *testing.T) {
	storage := NewURLStorage()
	key := "lol"
	badKey := "mek"
	value := "kek"
	storage.Add(key, value)

	type want struct {
		statusCode  int
		contentType string
		location    string
		body        string
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{name: "found", request: key, want: want{
			statusCode:  307,
			contentType: "",
			body:        "",
			location:    value,
		}},
		{name: "not found", request: badKey, want: want{
			statusCode:  404,
			contentType: "text/plain; charset=utf-8",
			body:        "url not found\n",
			location:    "",
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", nil)
			request.SetPathValue("id", test.request)
			w := httptest.NewRecorder()
			IDHandler(storage)(w, request)

			res := w.Result()

			assert.Equal(t, test.want.statusCode, res.StatusCode)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, test.want.location, res.Header.Get("Location"))
			assert.Equal(t, test.want.body, string(resBody))
		})
	}
}

func TestRootHandler(t *testing.T) {
	storage := NewURLStorage()

	type want struct {
		statusCode  int
		contentType string
		body        string
	}

	type request struct {
		method string
		id     string
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "get_request",
			request: request{
				method: http.MethodGet,
				id:     "kekovka",
			},
			want: want{
				statusCode:  http.StatusMethodNotAllowed,
				contentType: "text/plain; charset=utf-8",
				body:        "Only POST requests are allowed!\n",
			},
		},
		{
			name: "good_post_request",
			request: request{
				method: http.MethodPost,
				id:     "kekovka",
			},
			want: want{
				statusCode:  http.StatusCreated,
				contentType: "",
				body:        "http://example.com/" + doShort("kekovka"),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.request.method, "/", strings.NewReader(test.request.id))
			w := httptest.NewRecorder()
			RootHandler(storage)(w, request)

			res := w.Result()

			assert.Equal(t, test.want.statusCode, res.StatusCode)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, test.want.body, string(resBody))
		})
	}
}
