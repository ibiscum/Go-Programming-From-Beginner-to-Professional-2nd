package main

import (
	"testing"
)

func TestRemoveBad(t *testing.T) {
	if _, exists := getTypeName(""); exists {
		t.Fail()
	}
	if name, exists := getTypeName("305"); !exists || name != "Sue" {
		t.Fail()
	}
}
