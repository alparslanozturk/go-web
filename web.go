package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pg"
)

var DB_DSN = os.Getenv("DB_DSN")

const (
	DSN = "postgres://postgres:parola@localhost:5432/postgres "
	SQL = "select pg_is_in_recovery()"
)

func kokDizin(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Merhaba \n")
}
func pgCheck(w http.ResponseWriter, req *http.Request){
	
	db, err := sql.Open("postgres", DSN)
	if err != nil { panic(err) }

	rows, err := db.Query(SQL)
	if err != nil { panic(err) }

	for rows.Next(){
		switch err := rows.Scan(&)
	}


}



func main() {
	http.HandleFunc("/", kokDizin)
	http.ListenAndServe(":7777", nil)
}
