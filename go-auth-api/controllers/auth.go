package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "go-auth-api/config"
    "go-auth-api/models"
    "go-auth-api/utils"

    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    collection := config.GetUserCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    user.Password = string(hashedPassword)

    _, err := collection.InsertOne(ctx, user)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var loginData models.User
    _ = json.NewDecoder(r.Body).Decode(&loginData)

    collection := config.GetUserCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user models.User
    err := collection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    tokenString, _ := utils.GenerateJWT(user.Email)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
