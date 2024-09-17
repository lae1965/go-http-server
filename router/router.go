package router

import (
	"http-server/users"
	"net/http"
)

func Router() {
	http.HandleFunc("/user", users.UsersController)
	http.HandleFunc("/user/", users.UsersController)
}
