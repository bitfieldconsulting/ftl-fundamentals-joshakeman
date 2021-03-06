package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Howdy, Gopherinos!"
	got := hello.Greeting()
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
