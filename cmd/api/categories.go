package main

import (
	"log"
	"net/http"

	"github.com/igatto/todo/internal/store"
)

type CreateCategoryPayload struct {
	Name string `json:"name"`
}

func (app *application) createCategory(w http.ResponseWriter, r *http.Request) {
	var payload CreateCategoryPayload
	if err := readJSON(w, r, &payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	category := store.Category{
		Name: store.NewNullString(payload.Name),
	}
	ctx := r.Context()
	err := app.store.Categories.Create(ctx, category)
	if err != nil {
		log.Output(1, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
