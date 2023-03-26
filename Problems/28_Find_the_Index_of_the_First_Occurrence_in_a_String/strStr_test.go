package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name         string
		got_haystack string
		got_needle   string
		want         int
	}{
		{
			name:         "first case",
			got_haystack: "sadbutsad",
			got_needle:   "sad",
			want:         0,
		},
		{
			name:         "second case",
			got_haystack: "leetcode",
			got_needle:   "leeto",
			want:         -1,
		},
		{
			name:         "third case",
			got_haystack: "a",
			got_needle:   "a",
			want:         0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, strStr(test.got_haystack, test.got_needle))
		})
	}
}
