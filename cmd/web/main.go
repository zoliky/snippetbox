package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Dependency injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// Use the flag.Parse() function to parse the command-line fllag.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
	}

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	
	// Use the http.ListenAndServe() function to start a new web server.
	// If http.ListenAndServe() returns an error, use the log.Fatal() function
	// to log the error message and exit
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
