package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Initialize a new servemux, and register the home function as the handler
	// Make sure to use NewServeMux to prevent any security issues with the webapp
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}