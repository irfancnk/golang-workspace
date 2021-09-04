package main

import (
    "fmt"
    "net/http"
)

func getRoute(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Potato \n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", getRoute)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}
