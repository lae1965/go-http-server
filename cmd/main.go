package main

import (
	"fmt"
	"http-server/router"
	"log"
	"net/http"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	const PORT string = "8080"

	router.Router()

	fmt.Printf("Server has been started on %s Port...", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
