package routes

import (
    "github.com/gorilla/mux"
    "go-auth-api/controllers"
)

func RegisterAuthRoutes(r *mux.Router) {
    r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
    r.HandleFunc("/login", controllers.Login).Methods("POST")
}
