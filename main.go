package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "todo-app/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
    r.HandleFunc("/add", handlers.AddTaskHandler).Methods("POST")

    // Serve static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", r)
}
