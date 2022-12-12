package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2022, 12, 13, 12, 13, 14, 0, time.UTC)
	got := humanDate(tm)
	want := "13 Dec 2022 at 12:13"
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
