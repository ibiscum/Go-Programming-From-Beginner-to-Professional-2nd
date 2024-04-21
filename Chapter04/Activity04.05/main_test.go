package main

import (
	"testing"
)

func TestRemoveBad(t *testing.T) {
	if exists := getLocales(); len(exists) == 0 {
		t.Fail()
	}
}
