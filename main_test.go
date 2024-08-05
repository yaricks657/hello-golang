package main

import "testing"

func TestHello(t *testing.T) {
	//sd
	want := "Hello Go"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
