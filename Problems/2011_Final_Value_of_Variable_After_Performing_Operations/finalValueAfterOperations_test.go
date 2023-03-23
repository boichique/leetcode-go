package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFinalValueAfterOperations(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want int
	}{
		{
			name: "first case",
			got:  []string{"--X", "X++", "X++"},
			want: 1,
		},
		{
			name: "second case",
			got:  []string{"++X", "++X", "X++"},
			want: 3,
		},
		{
			name: "third case",
			got:  []string{"X++", "++X", "--X", "X--"},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, finalValueAfterOperations(test.got))
		})
	}
}
