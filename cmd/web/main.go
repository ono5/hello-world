package main

import (
	"fmt"
	"helloworld/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Printf("Starting appllication on port %s\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
