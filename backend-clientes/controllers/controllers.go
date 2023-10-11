package controllers

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Cliente(w http.ResponseWriter, r *http.Request) {
	clients := []Client{
		{1, "Cliente 1"},
		{2, "Cliente 2"},
		{3, "Cliente 3"},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
