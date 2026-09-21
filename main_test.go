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
