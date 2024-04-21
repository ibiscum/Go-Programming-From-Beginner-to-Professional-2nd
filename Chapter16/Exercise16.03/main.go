package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	name, ok := vl["name"]
	if !ok {
		w.WriteHeader(400)
		_, err := w.Write([]byte("Missing name"))
		if err != nil {
			log.Panic(err)
		}
		return
	}
	_, err := w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
