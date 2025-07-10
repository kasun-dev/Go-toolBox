package main

import (
    "log"
    "net/http"

    "go-auth-api/config"
    "go-auth-api/routes"

    "github.com/gorilla/mux"
)

func main() {
    config.ConnectDB()

    r := mux.NewRouter()
    routes.RegisterAuthRoutes(r)

    log.Println("🚀 Server started on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
