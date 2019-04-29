package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatoaguimaraes/golang-microservice/model"
	"github.com/renatoaguimaraes/golang-microservice/repository"
)

// GetUsers return all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repository.GetUsers()
	json.NewEncoder(w).Encode(users)
}

// GetUser return a user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, _ := repository.GetUser(id)
	if user.IsValid() {
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// CreateUser crates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := model.User{ID: id}
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsValid() {
		repository.CreateUser(user)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeleteUser deletes a user specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := repository.GetUser(id)
	if err == nil && user.IsValid() {
		repository.DeleteUser(user.ID)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
