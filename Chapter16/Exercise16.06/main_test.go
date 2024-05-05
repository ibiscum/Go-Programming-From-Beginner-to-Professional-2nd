package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_something(t *testing.T) {
	srv := httptest.NewServer(http.StripPrefix(
		"/statics/",
		http.FileServer(http.Dir("./statics")),
	))

	rsp, err := http.Get(srv.URL + "/statics/body.css")
	if err != nil {
		t.Error(err)
	}

	expected, err := os.ReadFile("./statics/body.css")
	if err != nil {
		t.Error(err)
	}

	actual := make([]byte, rsp.ContentLength)
	_, err = rsp.Body.Read(actual)
	if err != io.EOF {
		panic(err)
	}
	if string(actual) != string(expected) {
		t.Errorf("%s\n%s", string(expected), string(actual))
	}
}
