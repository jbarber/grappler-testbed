package main

import "testing"

func TestGreet(t *testing.T) {
	if got := Greet("world"); got != "hello, world" {
		t.Errorf("Greet = %q", got)
	}
}
