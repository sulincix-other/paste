package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", view)
    http.HandleFunc("/paste", index_html)
    http.HandleFunc("/api", paste)
    http.ListenAndServe(":2023", nil)
}
