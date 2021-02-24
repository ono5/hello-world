package main

import (
	"fmt"
	"log"
	"net/http"

	"helloworld/pkg/config"
	"helloworld/pkg/handlers"
	"helloworld/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers((repo))

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Printf("Starting appllication on port %s\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
