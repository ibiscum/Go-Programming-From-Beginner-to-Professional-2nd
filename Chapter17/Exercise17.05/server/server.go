package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check Authorization header is what we expect
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Authorization token not recognized"))
		if err != nil {
			log.Panic(err)
		}
		return
	}

	// wait 10 seconds before responding
	time.Sleep(10 * time.Second)

	msg := "hello client!"
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}

}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
