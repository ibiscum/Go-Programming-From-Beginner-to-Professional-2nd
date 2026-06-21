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

func TestMain_PrintsDebugAndLogLevel(t *testing.T) {
	out := captureStdout(t, main)
	out = strings.TrimSpace(out)
	parts := strings.Fields(out)
	if len(parts) < 2 {
		t.Fatalf("unexpected output format: %q", out)
	}

	if parts[0] != "false" {
		t.Fatalf("expected first field to be false, got %q", parts[0])
	}
	if parts[1] != "info" {
		t.Fatalf("expected second field to be info, got %q", parts[1])
	}
}
