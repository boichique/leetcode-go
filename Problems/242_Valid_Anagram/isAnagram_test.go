package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		got1 string
		got2 string
		want bool
	}{
		{
			name: "true case",
			got1: "anagram",
			got2: "nagaram",
			want: true,
		},
		{
			name: "false case",
			got1: "rat",
			got2: "car",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, isAnagram(test.got1, test.got2))
		})
	}
}
