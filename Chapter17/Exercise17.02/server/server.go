package main

import (
	"log"
	"net/http"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "{\"message\": \"hello world\"}"
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
