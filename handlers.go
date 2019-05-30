package main

import (
	"encoding/json"
	"net/http"
)

type echoResponse struct {
	Message string `json:"message"`
}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	resp := echoResponse{Message: message}

	respStr, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respStr)
}
