package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}

type messageData struct {
	Message string `json:"message"`
}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)

	messageData := messageData{}
	err := jsonDecoder.Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, _ := json.Marshal(messageData)
	log.Println(string(jsonBytes))
	_, err = w.Write(jsonBytes)
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
