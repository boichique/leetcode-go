package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want int
	}{
		{
			name: "first case",
			got:  "Hello World",
			want: 5,
		},
		{
			name: "second case",
			got:  "   fly me   to   the moon  ",
			want: 4,
		},
		{
			name: "third case",
			got:  "luffy is still joyboy",
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, lengthOfLastWord(test.got))
		})
	}
}
