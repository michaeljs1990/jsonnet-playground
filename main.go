package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	store PersistJsonnet
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	memFlag := flag.Bool("in-memory", false, "start up the daemon with an in-memory db (test only)")
	sqlFlag := flag.Bool("sql", false, "start up the daemon with a sql db (JSONNET_MYSQL_CONN) env for conn string")

	flag.Parse()

	// Add in toggle for the different types of backends
	if *memFlag {
		store = InMemory{
			store: map[string]string{},
		}
	}

	if *sqlFlag {
		// EXAMPLE: "user:password@/dbname"
		db, err := sql.Open("mysql", os.Getenv("JSONNET_MYSQL_CONN"))
		if err != nil {
			panic(err)
		}
		store = NewJSQL(db)
	}

	server := &http.Server{
		Handler: http.TimeoutHandler(http.DefaultServeMux, 7*time.Second, ""),
		Addr:    ":" + port,
	}

	http.HandleFunc("/", HandleEditor)
	http.HandleFunc("/backend/share", HandleShare)
	http.HandleFunc("/backend/compile", HandleCompile)

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatalf("Error listening on :%v: %v", port, server.ListenAndServe())
}
