package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	if got := Greet("world"); got != "hello, world" {
		t.Errorf("Greet = %q", got)
	}
}

func TestRunPlain(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, false); err != nil {
		t.Fatalf("run: %v", err)
	}
	if got := out.String(); got != "hello, world\n" {
		t.Errorf("run = %q", got)
	}
}

func TestRunJSON(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, true); err != nil {
		t.Fatalf("run: %v", err)
	}
	if got := out.String(); got != "{\"greeting\":\"hello, world\"}\n" {
		t.Errorf("run = %q", got)
	}
}

func TestGreetingWriterToStderr(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if err := run(greetingWriter(&stdout, &stderr, true), false); err != nil {
		t.Fatalf("run: %v", err)
	}
	if got := stderr.String(); got != "hello, world\n" {
		t.Errorf("stderr = %q", got)
	}
	if got := stdout.String(); got != "" {
		t.Errorf("stdout = %q, want nothing on stdout", got)
	}
}

func TestGreetingWriterDefaultsToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if err := run(greetingWriter(&stdout, &stderr, false), false); err != nil {
		t.Fatalf("run: %v", err)
	}
	if got := stdout.String(); got != "hello, world\n" {
		t.Errorf("stdout = %q", got)
	}
	if got := stderr.String(); got != "" {
		t.Errorf("stderr = %q, want nothing on stderr", got)
	}
}
