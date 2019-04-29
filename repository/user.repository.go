package repository

import (
	"log"

	"github.com/renatoaguimaraes/golang-microservice/model"
	"github.com/renatoaguimaraes/golang-microservice/util"
)

// GetUsers get all users
func GetUsers() []model.User {
	iter := util.Session().Query("SELECT id, firstname, lastname FROM users").Iter()
	var users []model.User
	var user model.User
	for iter.Scan(&user.ID, &user.FirstName, &user.LastName) {
		users = append(users, user)
	}
	return users
}

// GetUser get by id
func GetUser(id string) (model.User, error) {
	var user model.User
	if err := util.Session().Query("SELECT id, firstname, lastname FROM users WHERE id = ? LIMIT 1", id).Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser create new user
func CreateUser(user model.User) error {
	if err := util.Session().Query("INSERT INTO users (id, firstname, lastname) VALUES (?, ?, ?)", user.ID, user.FirstName, user.LastName).Exec(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteUser delete by id
func DeleteUser(id string) error {
	if err := util.Session().Query("DELETE FROM users WHERE id = ?", id).Exec(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
