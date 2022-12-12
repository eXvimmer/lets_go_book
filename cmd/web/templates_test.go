package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 9, 26, 10, 11, 12, 0, time.UTC),
			want: "26 Sep 2022 at 10:11",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET", // Central European Time
			tm:   time.Date(2022, 9, 26, 10, 11, 12, 0, time.FixedZone("CET", 1*60*60)),
			want: "26 Sep 2022 at 09:11",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := humanDate(test.tm); got != test.want {
				t.Errorf("got: %q, want: %q", got, test.want)
			}
		})
	}
}
