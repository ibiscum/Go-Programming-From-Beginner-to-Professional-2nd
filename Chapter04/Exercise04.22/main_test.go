package main

import (
	"testing"
)

func TestConversion(t *testing.T) {
	// m := `
	// int8    = 127  > int64    = 127
	// int     = 128  > int8     = -128
	// int8    = 127  > float32 = 127
	// float64 = 3.14 > int     = 3`

	if _, err := doubler("bool"); err != nil {
		t.Fail()
	}
}
