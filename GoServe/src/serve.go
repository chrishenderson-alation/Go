package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is my response\n"))
	})
	http.HandleFunc("/anything[0-5]/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is my other response\n"))
	})
	http.ListenAndServe(":8080", nil)
}
