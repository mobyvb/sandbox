go
    package main

    import (
      "fmt"
      "math/rand"
      "net/http"
      "time"
    )

    func main() {
      http.HandleFunc("/", serveIndex)
      http.HandleFunc("/api/number", getNumber)
      fmt.Println("Listening on port 8080...")
      http.ListenAndServe(":8080", nil)
    }

    func serveIndex(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w, r, "index.html")
    }

    func getNumber(w http.ResponseWriter, r *http.Request) {
      rand.Seed(time.Now().UnixNano())
      number := rand.Intn(23) + 12
      fmt.Fprintf(w, "%d", number)
    }