package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api/number", func(w http.ResponseWriter, r *http.Request) {
		randomNumber := rand.Intn(31) + 20
		fmt.Fprint(w, strconv.Itoa(randomNumber))
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
