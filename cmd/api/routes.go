package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Config) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/inventory", InventoryHandler)
	return r
}
