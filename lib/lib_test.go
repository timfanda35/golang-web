package lib

import "testing"

func TestIndexTitle(t *testing.T) {
	param := "hello, world"
	want := "HELLO, WORLD"
	if got := IndexTitle(param); got != want {
		t.Errorf("IndexTitle() = %q, want %q", got, want)
	}
}
