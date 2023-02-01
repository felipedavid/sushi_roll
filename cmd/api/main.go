package main

import (
    "log"
    "flag"
    "os"
    "net/http"
    "time"
)

const version = "1.0.0"

type config struct {
    addr string
    env string
}

type app struct {
    config config
    logger *log.Logger
}

func main() {
    var cfg config
    flag.StringVar(&cfg.addr, "addr", "127.0.0.1:8000", "Server address")
    flag.StringVar(&cfg.env, "env", "dev", "dev|staging|prod")
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

    a := &app{
        config: cfg,
        logger: logger,
    }

    srv := &http.Server{
        Addr: cfg.addr,
        Handler: a.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    logger.Printf("starting %s server on %s", cfg.env, cfg.addr)
    err := srv.ListenAndServe()
    logger.Fatal(err.Error())
}
