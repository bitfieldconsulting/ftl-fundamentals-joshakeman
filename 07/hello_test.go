package hello_test

import (
	"hello"
	"testing"
)

func TestReturnGreeting(t *testing.T) {
	want := "Hi there yourself!"
	got := hello.ReturnGreeting("Hi there")

	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
