package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdoutAndPanic(t *testing.T, fn func()) (string, any) {
	t.Helper()

	oldStdout := os.Stdout
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	os.Stdout = writePipe

	var recovered any
	func() {
		defer func() {
			recovered = recover()
		}()
		fn()
	}()

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

	return string(output), recovered
}

func TestMain_PrintsExpectedOutputBeforePanic(t *testing.T) {
	got, _ := captureStdoutAndPanic(t, main)
	want := "5\nПривет, мир\n"

	if got != want {
		t.Fatalf("unexpected output before panic\nwant: %q\ngot:  %q", want, got)
	}
}

func TestMain_PanicsOnOutOfRangeAccess(t *testing.T) {
	_, recovered := captureStdoutAndPanic(t, main)
	if recovered == nil {
		t.Fatal("expected panic, got nil")
	}

	msg := fmt.Sprint(recovered)
	if !strings.Contains(msg, "index out of range") {
		t.Fatalf("expected panic message to contain %q, got %q", "index out of range", msg)
	}
}
