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

	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}
	os.Stdout = originalStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured output: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("failed to close reader: %v", err)
	}

	return buf.String()
}

func TestMain_PrintsExpectedOutputWithDefaultOffset(t *testing.T) {
	original := defaultOffset
	t.Cleanup(func() { defaultOffset = original })

	defaultOffset = 10

	output := captureStdout(t, main)
	gotLines := strings.Split(strings.TrimSpace(output), "\n")
	wantLines := []string{"10", "20"}

	if len(gotLines) != len(wantLines) {
		t.Fatalf("unexpected number of lines: got %d, want %d (output=%q)", len(gotLines), len(wantLines), output)
	}

	for i := range wantLines {
		if gotLines[i] != wantLines[i] {
			t.Fatalf("unexpected line %d: got %q, want %q", i+1, gotLines[i], wantLines[i])
		}
	}
}

func TestMain_PrintsExpectedOutputWithCustomOffset(t *testing.T) {
	original := defaultOffset
	t.Cleanup(func() { defaultOffset = original })

	defaultOffset = 7

	output := captureStdout(t, main)
	gotLines := strings.Split(strings.TrimSpace(output), "\n")
	wantLines := []string{"7", "14"}

	if len(gotLines) != len(wantLines) {
		t.Fatalf("unexpected number of lines: got %d, want %d (output=%q)", len(gotLines), len(wantLines), output)
	}

	for i := range wantLines {
		if gotLines[i] != wantLines[i] {
			t.Fatalf("unexpected line %d: got %q, want %q", i+1, gotLines[i], wantLines[i])
		}
	}
}
