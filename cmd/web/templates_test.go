package main

import (
	"testing"
	"time"

	"github.com/exvimmer/lets_go/snippetbox/internal/assert"
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
			assert.Equal(t, humanDate(test.tm), test.want)
		})
	}
}
