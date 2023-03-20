package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// We should be generating the version dynamically at build time
const version = "1.0.0"

type config struct {
	port int
	env  string
}

// application stores the dependencies for our handlers, helpers and middlewares
type application struct {
	config     config
	errLogger  *log.Logger
	infoLogger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Port for the HTTP server listen to")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	errLogger := log.New(os.Stderr, "[ERROR] ", log.Lshortfile|log.Ldate|log.Ltime)
	infoLogger := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	app := &application{
		config:     cfg,
		errLogger:  errLogger,
		infoLogger: infoLogger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%d", cfg.port),
		Handler:      mux,
		ErrorLog:     app.errLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLogger.Printf("Starting %s server at %s", cfg.env, server.Addr)
	err := server.ListenAndServe()
	errLogger.Fatal(err)
}
