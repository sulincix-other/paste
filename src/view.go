package main

import (
    "net/http"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func view (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TODO: implement this")
}
