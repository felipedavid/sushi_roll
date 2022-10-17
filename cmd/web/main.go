package main

import (
	"database/sql"
	"flag"
	"github.com/felipedavid/sushi_roll/internal/models"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// app contém maior parte do estado necessário para a operação da aplicação
type app struct {
	infoLog       *log.Logger
	errLog        *log.Logger
	game          *models.GameModel
	templateCache templateCache
	env           string
}

func main() {
	// Fazendo parsing dos argumentos por linha de comando
	addr := flag.String("addr", ":4000", "HTTP listen address")
	dsn := flag.String("dsn",
		"postgres://postgres:postgres@localhost/sushi_roll_db?sslmode=disable",
		"Database Service Name")
	env := *flag.String("env", "development", "Environment (development|production)")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err.Error())
	}

	tc, err := newTemplateCache()
	if err != nil {
		errLog.Fatal(err.Error())
	}

	a := app{
		infoLog:       infoLog,
		errLog:        errLog,
		game:          &models.GameModel{DB: db},
		templateCache: tc,
		env:           env,
	}

	infoLog.Printf("Starting server on address %s\n", *addr)
	err = http.ListenAndServe(*addr, a.routes())
	errLog.Fatal(err)
}

// openDB cria uma "connection pool" e testa se é possível se conectar ao banco de dados
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
