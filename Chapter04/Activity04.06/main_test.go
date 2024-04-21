package main

import (
	"testing"
)

func TestRemoveBad(t *testing.T) {
	if exists := getTypeName("int"); exists == "int" {
		t.Fail()
	}
}
