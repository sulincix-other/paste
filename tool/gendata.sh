echo 'package main'
echo ''
echo 'import ('
echo '    "fmt"'
echo '    "net/http"'
echo ')'
echo ''
echo 'func '$1' (w http.ResponseWriter, req *http.Request) {'
echo '    fmt.Fprintf(w, `'
cat "$2"
echo ''
echo '     `)'
echo '}'
