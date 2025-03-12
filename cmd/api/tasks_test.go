package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igatto/todo/internal/store"
)

func TestGetAllTasks(t *testing.T) {
	store := store.NewMockStorage()
	app := &application{
		store: store,
	}
	mux := app.mount()

	t.Run("request should return status OK", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
