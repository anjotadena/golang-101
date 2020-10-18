package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User handler"))
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Project Handler"))
}

func main() {
	fmt.Println("Welcome to my HTTP Server")

	// Handlers
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/users/", userHandler)
	http.HandleFunc("/projects/", projectsHandler)

	http.ListenAndServe(":8000", nil)
}
