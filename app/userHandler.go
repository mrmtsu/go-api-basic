package app

import (
	"encoding/json"
	"go-api-basic/domain"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := []domain.User{}
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]

	user := domain.User{}
	DB.First(&user, userId)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	updateUser := domain.User{}
	json.NewDecoder(r.Body).Decode(&updateUser)
	DB.Save(&updateUser)

	json.NewEncoder(w).Encode(updateUser)
}
