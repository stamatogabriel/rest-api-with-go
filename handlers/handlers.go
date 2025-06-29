package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/services"
)

var todo services.Todo

// HealthCheckHandler is a simple handler to check if the API is running
func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	res := Response{
		Msg:  "API is running",
		Code: 200,
	}

	jsonResponse, err := json.Marshal(res)

	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
