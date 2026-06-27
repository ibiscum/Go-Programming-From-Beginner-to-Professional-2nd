package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe create: %v", err)
	}
	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}
	os.Stdout = orig

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("read pipe: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("close reader: %v", err)
	}
	return buf.String()
}

func TestMain_PrintsBetween1And5Stars(t *testing.T) {
	out := captureStdout(t, main)
	s := strings.TrimSpace(out)
	if s == "" {
		t.Fatalf("no output captured")
	}
	for _, r := range s {
		if r != '*' {
			t.Fatalf("unexpected rune in output: %q", r)
		}
	}
	if n := len(s); n < 1 || n > 5 {
		t.Fatalf("unexpected number of stars: got=%d want between 1 and 5", n)
	}
}
