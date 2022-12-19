package main

import (
    "net/http"
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func view (w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s\n",r.URL.Path)
    if r.URL.Path == "/" {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, "<meta http-equiv=\"Refresh\" content=\"0; url='/paste'\" />")
        return
    }
    db, err := sql.Open("sqlite3", "./paste.db")
    if err != nil {
        fmt.Println(err)
    }

    var paste string
    query := fmt.Sprintf("SELECT paste FROM paste WHERE id='%s';", r.URL.Path[1:])
    err = db.QueryRow(query).Scan(&paste)

    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Fprintf(w, "%s\n", paste)
}
