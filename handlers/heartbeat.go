package handlers

import (
	"encoding/json"
	"net/http"
)

type HeartbeatResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Heartbeat(rw http.ResponseWriter, r *http.Request) {
	response := HeartbeatResponse{
		Status:  http.StatusOK,
		Message: "Online...",
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}
