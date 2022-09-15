package main

import "testing"

func TestSumm(t *testing.T) {
	got := summOfDigits(256, 42)
	want := 298

	if got != int64(want) {
		t.Errorf("got %q want %q", got, want)
	}
}
