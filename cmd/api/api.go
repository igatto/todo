package main

import (
	"fmt"
	"net/http"

	"github.com/igatto/todo/internal/store"
)

type application struct {
	store store.Storage
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", app.getAllTasks)
	mux.HandleFunc("POST /categories", app.createCategory)
	return mux
}

func (app *application) run(mux *http.ServeMux) {
	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", mux)
}
