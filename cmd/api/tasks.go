package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (app *application) getAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := app.store.Tasks.GetAll(ctx)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	j, err := json.Marshal(tasks)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
