package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the home page handler
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Printf("Starting appllication on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
