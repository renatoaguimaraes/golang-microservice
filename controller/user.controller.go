package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatoaguimaraes/golang-microservice/model"
)

var users []model.User

func init() {
	users = append(users, model.User{ID: "renatoaguimaraes@gmail.com", FirstName: "Renato", LastName: "Guimarães"})
}

// GetUsers return all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// GetUser return a user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.User{})
}

// CreateUser crates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser deletes a user specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(users)
	}
}
