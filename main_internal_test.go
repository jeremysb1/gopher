package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "hello world"

	got := greet(lang)

	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "bonjour le monde"

	got := greet(lang)

	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}