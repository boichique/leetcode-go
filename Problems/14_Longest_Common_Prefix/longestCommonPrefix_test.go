package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want string
	}{
		{
			name: "base case",
			got:  []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "empty string case",
			got:  []string{"", "city", "country"},
			want: "",
		},
		{
			name: "second case",
			got:  []string{"a"},
			want: "a",
		},
		{
			name: "third case",
			got:  []string{"car", "cir"},
			want: "c",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, longestCommonPrefix(test.got))
		})
	}
}
