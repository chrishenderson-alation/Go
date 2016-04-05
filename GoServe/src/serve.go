package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("This is my response"))
    })
    http.ListenAndServe(":8080", nil)
    // http.ListenAndServeTLS(addr, certFile, keyFile, handler)
    http.
}
