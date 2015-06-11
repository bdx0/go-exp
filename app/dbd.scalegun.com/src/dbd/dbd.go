package main

import (
    "scalegun.com/v1"
    "net/http"
    "fmt"
)

func first(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "first!")    
}

func main () {
    app := scalegun.NewApp("dbd.scalegun.com", "secret")
    app.HandleFunc("/", first)
    app.ListenAndServe()
}