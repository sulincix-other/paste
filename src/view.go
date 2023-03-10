package main

import (
    "net/http"
    "fmt"
    "strings"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func view (w http.ResponseWriter, r *http.Request) {
    if strings.Contains(r.UserAgent(),"Windows") || strings.Contains(r.UserAgent(),"iPhone") {
        return
    }

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
    query := fmt.Sprintf("SELECT paste FROM paste WHERE id='%s';", b64_encode(r.URL.Path[1:]))
    err = db.QueryRow(query).Scan(&paste)

    if err != nil {
        fmt.Println(err)
    }
    w.Header().Set("Content-type:", "text/plain; charset=utf-8");
    fmt.Fprintf(w, "%s", b64_decode(paste))
}
