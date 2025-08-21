package v2

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Shivansh")
	want := "Hello Shivansh"
	if got == want {
		t.Errorf("got %q want %q", got, want)
	}

}
