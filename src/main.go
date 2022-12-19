package main

import (
    "net/http"
)

func main() {

    http.HandleFunc("/", index_html)
    http.HandleFunc("/paste", paste)
    http.HandleFunc("/view", view)
    http.ListenAndServe(":8090", nil)
}
