package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestStatic(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	http.ServeFile(w, r, "./index.html")

	content, err := os.ReadFile("index.html")
	if err != nil {
		t.Error(err)
	}
	//rpl := strings.NewReplacer(" ","","\n","")
	if w.Body.String() != string(content) {
		t.Errorf("%s\n%s", string(content), w.Body.String())
	}
}
