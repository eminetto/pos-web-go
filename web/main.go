package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/eminetto/pos-web-go/core/beer"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("sqlite3", "data/beer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service := beer.NewService(db)
	/*mux := http.NewServeMux()
	mux.Handle("/", hello(service))*/
	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.NewLogger(),
	)
	r.Handle("/v1/beer", n.With(
		negroni.Wrap(hello(service)),
	)).Methods("GET", "OPTIONS")
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		//Handler:      mux,
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func hello(service *beer.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		all, _ := service.GetAll()
		fmt.Println(all)
	})
}