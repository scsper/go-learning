package main

import (
    "net/http"
    "fmt"
)

type String string
type Struct struct {
    Greeting string
    Punct string
    Who String
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s.Greeting)
    fmt.Fprint(w, s.Punct)
    fmt.Fprint(w, s.Who)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

    http.ListenAndServe("localhost:4000", nil)
}
