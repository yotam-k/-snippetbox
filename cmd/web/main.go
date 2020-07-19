package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// Defines command line flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	// Parses command line arguments for the above defined flags
	flag.Parse()

	// Logger for informational messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Logger for error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	// Initialize a new servemux, and register the home function as the handler
	// Make sure to use NewServeMux to prevent any security issues with the webapp
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Create our own server to use the errorLog defined above
	// http.ListenAndServe() resorts to the standard logger, not a custom logger
	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}