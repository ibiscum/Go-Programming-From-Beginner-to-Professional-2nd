package main

import (
	"testing"
)

func TestRemoveBad(t *testing.T) {
	if _, exists := getLocales(""); exists {
		t.Fail()
	}
	if name, exists := getLocales("305"); !exists || name != "Sue" {
		t.Fail()
	}
}
