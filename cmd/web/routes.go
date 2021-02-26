package main

import (
	"net/http"

	"helloworld/pkg/config"
	"helloworld/pkg/handlers"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) *pat.PatternServeMux {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
