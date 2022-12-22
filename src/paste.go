package main

import (
    "net/http"
    "fmt"
    "database/sql"
    "os"
    "strings"
    "hash/crc32"
    _ "github.com/mattn/go-sqlite3"
)

func paste (w http.ResponseWriter, r *http.Request) {
    // variables
    initdb := false
    web := "0"
    var paste string
    var paste_id string

    if strings.Contains(r.UserAgent(),"Windows") || strings.Contains(r.UserAgent(),"iPhone") {
        return
    }
    // Headers
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    // Request check
    if "POST" == r.Method {
        web = r.FormValue("web")
        paste = r.FormValue("paste")
    }else{
        fmt.Fprintf(w,"Invalid request type: %s", r.Method)
    }

    // DB check
    if _, err := os.Stat("./paste.db"); err != nil {
        initdb = true
    }
    db, err := sql.Open("sqlite3", "./paste.db")
    if err != nil {
        fmt.Println(err)
    }
    if initdb {
        // Create table if required
        db.Exec("CREATE TABLE paste (id text, paste text)")
    }

    // Calculate crc32
    crc32q := crc32.MakeTable(0xD5828281)
    paste_id = fmt.Sprintf("%08x", crc32.Checksum([]byte(string(paste)), crc32q))

    // Already exists check
    var id_exists string
    query := fmt.Sprintf("SELECT id FROM paste WHERE id='%s';", b64_encode(paste_id))
    err = db.QueryRow(query).Scan(&id_exists)
    if err != nil {
        fmt.Println(err)
    }
    
    // Insert new paste
    if id_exists != b64_encode(paste_id) {
        query := fmt.Sprintf("INSERT OR REPLACE INTO paste (id,paste) VALUES(\"%s\", \"%s\");", b64_encode(paste_id), b64_encode(paste))
        _, err = db.Exec(query)
        if err != nil {
            fmt.Println(err)
        }
    }

    // Redirect
    if web == "1" {
        fmt.Fprintf(w, "<meta http-equiv=\"Refresh\" content=\"0; url='/%s'\" />", paste_id)
    }else{
        fmt.Fprintf(w, "%s\n", paste_id)
    }
}
