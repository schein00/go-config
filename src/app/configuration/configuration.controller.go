package configuration

import (
    "fmt"
    "net/http"
)

type Configuration struct {}


func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "config")
}

func (config *Configuration) Start() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}