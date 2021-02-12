package main

import (
    "database/sql" 
    "fmt"
    "net/http"

    _"github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "password"
    dbname = ""
)

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
  panic(err)
}

defer db.Close()
func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}

