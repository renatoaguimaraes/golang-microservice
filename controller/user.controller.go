package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/renatoaguimaraes/golang-microservice/model"
)

var session *gocql.Session

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "user"
	session, _ = cluster.CreateSession()
}

// GetUsers return all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	iter := session.Query("SELECT id, firstname, lastname FROM users").Iter()
	var users []model.User
	var user model.User
	for iter.Scan(&user.ID, &user.FirstName, &user.LastName) {
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// GetUser return a user by id
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// CreateUser crates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	if err := session.Query("INSERT INTO users (id, firstname, lastname) VALUES (?, ?, ?)", user.ID, user.FirstName, user.LastName).Exec(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}

// DeleteUser deletes a user specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
