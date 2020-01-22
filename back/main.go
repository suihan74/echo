package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func public(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello private!\n"))
}

func main() {
    allowedOrigins := handlers.AllowedOrigins([]string {"http://localhost:8080"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

    router := mux.NewRouter()
    router.HandleFunc("/public", public)
    router.HandleFunc("/private", private)

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
