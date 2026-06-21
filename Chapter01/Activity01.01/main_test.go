package main

import (
	"io"
	"os"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	os.Stdout = writePipe

	fn()

	if err := writePipe.Close(); err != nil {
		t.Fatalf("failed to close writer pipe: %v", err)
	}
	os.Stdout = oldStdout

	output, err := io.ReadAll(readPipe)
	if err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}
	if err := readPipe.Close(); err != nil {
		t.Fatalf("failed to close reader pipe: %v", err)
	}

	return string(output)
}

func TestMain_PrintsExpectedValues(t *testing.T) {
	got := captureStdout(t, main)

	want := "Bob\nSmith\n34\nfalse\n"
	if got != want {
		t.Fatalf("unexpected output from main()\nwant:\n%q\ngot:\n%q", want, got)
	}
}
