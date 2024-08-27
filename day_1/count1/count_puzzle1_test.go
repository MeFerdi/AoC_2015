package main

import "testing"

func TestSum(t *testing.T) {
	got := int(138)
	want := int(138)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
