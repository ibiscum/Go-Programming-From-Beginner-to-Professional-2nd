package main

import (
	"io"
	"os"
	"strings"
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
	out := captureStdout(t, main)
	out = strings.TrimSpace(out)
	lines := strings.Split(out, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines of output, got %d: %q", len(lines), out)
	}

	if lines[0] != "0 0 0" {
		t.Fatalf("unexpected first line\nwant: %q\ngot:  %q", "0 0 0", lines[0])
	}

	if lines[1] != "one 1 1.5 2 2.5" {
		t.Fatalf("unexpected second line\nwant: %q\ngot:  %q", "one 1 1.5 2 2.5", lines[1])
	}

	fields := strings.Fields(lines[2])
	if len(fields) < 2 {
		t.Fatalf("unexpected third line format: %q", lines[2])
	}
	if fields[0] != "false" {
		t.Fatalf("expected Debug to be false, got %q", fields[0])
	}
	if fields[1] != "info" {
		t.Fatalf("expected LogLevel to be info, got %q", fields[1])
	}
}
