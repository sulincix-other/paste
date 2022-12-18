package main

import (
    "net/http"
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func paste(w http.ResponseWriter, r *http.Request) {

    if "POST" == r.Method {
        paste := r.FormValue("paste")
        fmt.Fprintf(w, "%s\n", paste)
    }
    db, err := sql.Open("sqlite3", "./paste.db")
    if err != nil {
        fmt.Println(err)
    }
    var version string
    err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(version)
}
