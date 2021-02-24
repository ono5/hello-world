package main

import (
	"fmt"
	"helloworld/pkg/handlers"
	"log"
	"net/http"

	"github.com/ono5/hello-world/pkg/config"
	"github.com/ono5/helloworld/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Printf("Starting appllication on port %s\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
