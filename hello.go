package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])    
    
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        fmt.Fprintf(w, "Hello")
    }else{
       fmt.Fprintf(w, "Hello %s", name)
   }
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/query", queryHandler)
    http.ListenAndServe(":8080", nil)
}
