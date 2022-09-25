package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "zoli:dummYpass1%@/snippetbox?parseTime=true", "MySQL data source name")
	
	// Use the flag.Parse() function to parse the command-line fllag.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Pass openDB() the DSN from the command-line fllag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exists.
	defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
