package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumOfUnique(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want int
	}{
		{
			name: "first case",
			got:  []int{1, 2, 3, 2},
			want: 4,
		},
		{
			name: "second case",
			got:  []int{1, 1, 1, 1, 1},
			want: 0,
		},
		{
			name: "third case",
			got:  []int{1, 2, 3, 4, 5},
			want: 15,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, sumOfUnique(test.got))
		})
	}
}
