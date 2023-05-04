package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var number = 5

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/number", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			var request struct {
				Number int `json:"number"`
			}
			err := decoder.Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			number = request.Number
			fmt.Fprintf(w, "Number stored: %d", number)
		} else {
			response := struct {
				Number int `json:"number"`
			}{
				Number: number,
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
