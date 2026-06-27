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

func TestMain_PrintsExpectedSequence_DefaultLevel(t *testing.T) {
	output := captureStdout(t, main)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	want := []string{
		"Main start  : pkg",
		"Block start : pkg",
		"funcA start : pkg",
	}

	if len(lines) != len(want) {
		t.Fatalf("unexpected number of lines: got=%d want=%d output=%q", len(lines), len(want), output)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Fatalf("line %d: got=%q want=%q", i+1, lines[i], want[i])
		}
	}
}

func TestMain_PrintsExpectedSequence_ModifiedLevel(t *testing.T) {
	orig := level
	t.Cleanup(func() { level = orig })

	level = "test"
	output := captureStdout(t, main)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	want := []string{
		"Main start  : test",
		"Block start : test",
		"funcA start : test",
	}

	if len(lines) != len(want) {
		t.Fatalf("unexpected number of lines: got=%d want=%d output=%q", len(lines), len(want), output)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Fatalf("line %d: got=%q want=%q", i+1, lines[i], want[i])
		}
	}
}
