package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go server on Railway!")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // Default port if Railway doesn't set one
    }

    http.HandleFunc("/", handler)
    log.Printf("Server starting on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
