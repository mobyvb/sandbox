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

	http.HandleFunc("/api/numbers", func(w http.ResponseWriter, r *http.Request) {
		randomNumber := rand.Intn(24) + 14
		fmt.Fprint(w, strconv.Itoa(randomNumber))
	})

	fmt.Println("Serving on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
