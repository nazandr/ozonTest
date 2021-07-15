package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nazandr/ozonTest/internal/app/models"
	"github.com/nazandr/ozonTest/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestServer_handleShort(t *testing.T) {
	store, teardown := store.TestStore(t, "host=localhost dbname=short_url_test sslmode=disable")
	defer teardown("urls")
	server := TestServer(t)
	server.Store = store

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"url": "example.com/long",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid url",
			payload: map[string]string{
				"url": "invalid",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/short", b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.Echo.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleLong(t *testing.T) {
	store, teardown := store.TestStore(t, "host=localhost dbname=short_url_test sslmode=disable")
	defer teardown("urls")
	server := TestServer(t)
	server.Store = store

	url := models.NewURL()
	url.Long = "example.com/long"
	if err := server.Store.Url().Create(url); err != nil {
		t.Fatal(err)
	}

	url.Shortener()
	server.Store.Url().UpdateShort(url)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"url": url.Short,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid url",
			payload: map[string]string{
				"url": "invalid",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/long", b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.Echo.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
