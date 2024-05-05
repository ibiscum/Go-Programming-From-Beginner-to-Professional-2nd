package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_name(t *testing.T) {
	hdl, err := NewHello("./index.html")
	if err != nil {
		t.Error(err)
	}
	srv := httptest.NewServer(hdl)
	rsp, err := http.Get(srv.URL + "/?name=john")
	if err != nil {
		t.Error(err)
	}
	expected, err := os.ReadFile("./teststatics/john.html")
	if err != nil {
		t.Error(err)
	}
	actual := make([]byte, rsp.ContentLength)
	_, err = rsp.Body.Read(actual)
	if err != io.EOF {
		log.Panic(err)
	}
	if string(actual) != string(expected) {
		t.Errorf("\n%s\n%s", string(expected), string(actual))
	}
}

func Test_anonymous(t *testing.T) {
	hdl, err := NewHello("./index.html")
	if err != nil {
		t.Error(err)
	}
	srv := httptest.NewServer(hdl)
	rsp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Error(err)
	}
	expected, err := os.ReadFile("./teststatics/anonymous.html")
	if err != nil {
		t.Error(err)
	}
	actual := make([]byte, rsp.ContentLength)
	_, err = rsp.Body.Read(actual)
	if err != io.EOF {
		log.Panic(err)
	}
	if string(actual) != string(expected) {
		t.Errorf("\n%s\n%s", string(expected), string(actual))
	}
}
