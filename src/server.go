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

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World")
}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
    panic(err)
    }

    defer db.Close()

    err = db.Ping()
    if err != nil {
    panic(err)
    }

    fmt.Println("Successfully connected!")
    
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}

