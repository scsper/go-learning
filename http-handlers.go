package main

import (
    "net/http"
    "fmt"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}


func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))

    http.ListenAndServe("localhost:4000", nil)
}
