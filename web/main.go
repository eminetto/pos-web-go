package main

import (
	"database/sql"
	"github.com/codegangsta/negroni"
	"github.com/eminetto/pos-web-go/core/beer"
	"github.com/eminetto/pos-web-go/web/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "data/beer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service := beer.NewService(db)
	r := mux.NewRouter()
	//middlewares - código que vai ser executado em todas as requests
	//aqui podemos colocar logs, inclusão e validação de cabeçalhos, etc
	n := negroni.New(
		negroni.NewLogger(),
	)
	//handlers
	handlers.MakeBeerHandlers(r, n, service)

	//static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	r.PathPrefix("/static/").Handler(n.With(
		negroni.Wrap(http.StripPrefix("/static/", fileServer)),
	)).Methods("GET", "OPTIONS")

	http.Handle("/", r)

	//usado para testes de performance do servidor usando o siege
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}