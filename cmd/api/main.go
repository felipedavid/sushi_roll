package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config holds all the configuration settings for our app
type config struct {
	port int
	env  string
}

// app holds all the dependecies for our handlers, helpers and middleware
type app struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Port that the server should listen to")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stag|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	a := &app{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%d", cfg.port),
		Handler:      a.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s\n", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
