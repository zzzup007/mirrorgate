package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := "3001"
    if len(os.Args) > 1 {
        port = os.Args[1]
    }
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "🔥 Backend %s says: Hello! You requested: %s", port, r.URL.Path)
    })
    
    fmt.Printf("🚀 Backend starting on port %s\n", port)
    http.ListenAndServe(":"+port, nil)
}
