package main

import (
    "net/http"
    "fmt"
    "database/sql"
    "os"
    "time"
    "hash/crc32"
    _ "github.com/mattn/go-sqlite3"
)

func paste (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    initdb := false
    now := time.Now()
    if _, err := os.Stat("./paste.db"); err != nil {
        initdb = true
    }
    var paste string
    var paste_id string
    if "POST" == r.Method {
        paste = r.FormValue("paste")
    }else{
        fmt.Fprintf(w,"Invalid request type: %s", r.Method)
    }
    db, err := sql.Open("sqlite3", "./paste.db")
    if err != nil {
        fmt.Println(err)
    }
    if initdb {
        db.Exec("CREATE TABLE paste (id text, paste text)")
    }
    crc32q := crc32.MakeTable(0xD5828281)
    paste_id = fmt.Sprintf("%08x", crc32.Checksum([]byte(string(now.UnixNano())), crc32q))
    query := fmt.Sprintf("INSERT INTO paste (id,paste) VALUES(\"%s\", \"%s\");", paste_id, paste)
    fmt.Fprintf(w,"%s\n",query)
    _, err = db.Exec(query)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, "<meta http-equiv=\"Refresh\" content=\"0; url='/view/%s'\" />", paste_id)
}
