package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func rootContext(w http.ResponseWriter, req *http.Request) {

	var pg_is_in_recovery string

	db, _ := sql.Open("postgres", "postgres://postgres:parola@localhost:5432/postgres?sslmode=disable")

	err := db.QueryRow("select pg_is_in_recovery()").Scan(&pg_is_in_recovery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Down(Query Error)\n")
		return
	}

	switch pg_is_in_recovery {
	case "false":
		fmt.Fprintf(w, "Primary\n")
	case "true":
		fmt.Fprintf(w, "Standby\n")
	default:
		fmt.Fprintf(w, "Down\n")
	}
}

func main() {
	http.HandleFunc("/", rootContext)

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}
