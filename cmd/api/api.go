package main

import (
	"encoding/json"
	"net/http"

	"github.com/igatto/todo/internal/store"
)

type application struct {
	store store.Storage
}

func (app *application) run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleRoot)
	http.ListenAndServe(":8080", mux)
}

func (app *application) handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := app.store.Tasks.GetAll(ctx)
	j, err := json.Marshal(tasks)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
