package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadedFile, uploadedFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	fileContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}
	defer uploadedFile.Close()

	err = os.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0o600)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write([]byte(fmt.Sprintf("%s Uploaded!", html.EscapeString(uploadedFileHeader.Filename))))
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
