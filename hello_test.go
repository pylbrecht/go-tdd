package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("pylbrecht")
	want := "Hello pylbrecht"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
