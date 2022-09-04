package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type app struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {
	// Parsing command line flags
	addr := flag.String("addr", ":4000", "HTTP listen address")
	dsn := flag.String("dsn", "postgres://sushi:roll@localhost/sushi_roll_db?sslmode=disable", "Database Service Name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	_, err := openDB(*dsn)
	if err != nil {
		errLog.Println(err.Error())
		return
	}

	a := app{
		infoLog: infoLog,
		errLog:  errLog,
	}

	infoLog.Printf("Starting server on address %s\n", *addr)
	err = http.ListenAndServe(*addr, a.routes())
	errLog.Fatal(err)
}

// openDB creates a connection pool and test the connection with the database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
