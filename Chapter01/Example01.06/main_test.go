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
	fields := strings.Fields(out)

	if len(fields) < 4 {
		t.Fatalf("expected at least 4 fields in output, got %d: %q", len(fields), out)
	}

	if fields[0] != "false" {
		t.Fatalf("expected first field to be false, got %q", fields[0])
	}

	if fields[1] != "info" {
		t.Fatalf("expected second field to be info, got %q", fields[1])
	}

	// fields[2] and fields[3] for time might be combined if space is used in time format.
	// But strings.Fields splits by whitespace.
	// Let's check the last field which should be _A1_Μείγμα ("")
	lastField := fields[len(fields)-1]
	if lastField != "" {
		t.Fatalf("expected last field to be , got %q", lastField)
	}
}
