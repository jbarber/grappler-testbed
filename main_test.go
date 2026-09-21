package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	if got := Greet("world"); got != "hello, world" {
		t.Errorf("Greet = %q", got)
	}
}

func TestRunVersionFlag(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{"--version"}, &out); code != 0 {
		t.Errorf("run(--version) exit code = %d, want 0", code)
	}
	if got := strings.TrimSpace(out.String()); got != version {
		t.Errorf("run(--version) printed %q, want %q", got, version)
	}
}

func TestRunWithoutFlagsGreets(t *testing.T) {
	var out bytes.Buffer
	if code := run(nil, &out); code != 0 {
		t.Errorf("run() exit code = %d, want 0", code)
	}
	if got := strings.TrimSpace(out.String()); got != "hello, world" {
		t.Errorf("run() printed %q, want %q", got, "hello, world")
	}
}

func TestDeliberatelyBroken(t *testing.T) {
	if Greet("x") != "this will not match" {
		t.Fatal("deliberate CI failure for the grappler test plan")
	}
}
