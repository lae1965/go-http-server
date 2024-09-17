package users

import (
	"time"
)

//! Здесь будут вызовы к базе данных

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func getUserById(id int64) *User {
	for i := range users {
		if users[i].Id == id {
			return &users[i]
		}
	}
	return nil
}

func getAllUsers() []User {
	return users
}

func addUser(user User) User {
	user.Id = time.Now().UnixMilli()
	users = append(users, user)
	return user
}

func updateUser(id int64, newUser User) *User {
	user := getUserById(id)
	if user == nil {
		return nil
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}
	return user
}

func removeUser(id int64) bool {
	for i := range users {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
