package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Time: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	port := ":8795"
	log.Println("Server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
