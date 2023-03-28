package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name       string
		gotPattern string
		gotS       string
		want       bool
	}{
		{
			name:       "first case",
			gotPattern: "abba",
			gotS:       "dog cat cat dog",
			want:       true,
		},
		{
			name:       "second case",
			gotPattern: "abba",
			gotS:       "dog cat cat fish",
			want:       false,
		},
		{
			name:       "third case",
			gotPattern: "aaaa",
			gotS:       "dog cat cat dog",
			want:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, wordPattern(test.gotPattern, test.gotS))
		})
	}
}
