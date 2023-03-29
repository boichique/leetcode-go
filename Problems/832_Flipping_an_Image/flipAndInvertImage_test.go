package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		name string
		got  [][]int
		want [][]int
	}{
		{
			name: "first case",
			got:  [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name: "second case",
			got:  [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, flipAndInvertImage(test.got))
		})
	}
}
