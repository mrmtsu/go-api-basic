package app

import (
	"encoding/json"
	"go-api-basic/domain"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := []*domain.User{
		&domain.User{
			Id:       1,
			Name:     "a",
			Email:    "a@a.com",
			Password: "a",
		},
		&domain.User{
			Id:       2,
			Name:     "b",
			Email:    "b@b.com",
			Password: "b",
		},
	}
	json.NewEncoder(w).Encode(users)
}
