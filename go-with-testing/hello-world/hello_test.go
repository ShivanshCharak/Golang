package main

import "testing"

func TestHelo(t *testing.T) {
	got := hello()
	want := "hello string"
	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
}
