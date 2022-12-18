package main

import (
    "net/http"
)

func main() {

    http.HandleFunc("/", index_html)
    http.HandleFunc("/main.css", main_css)
    http.ListenAndServe(":8090", nil)
}
