package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/get-number", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(28) + 23 // to get number between 23 and 50

		fmt.Fprintf(w, `{"number": %s}`, strconv.Itoa(num))
	})

	http.ListenAndServe(":3000", nil)
}
