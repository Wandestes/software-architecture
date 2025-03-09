package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	port := ":8795"
	println("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
