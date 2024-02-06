package main

import (
	"msg-inventory/data"
	"net/http"
)

type InventoryPayLoad struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Data   string `json:"data"`
}

func (app *Config) InventoryHandler(w http.ResponseWriter, r *http.Request) {
	app.WriteJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "InventoryHandler",
		Data:    data.Inventory{},
	})
}
