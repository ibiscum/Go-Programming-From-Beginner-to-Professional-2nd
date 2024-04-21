package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
