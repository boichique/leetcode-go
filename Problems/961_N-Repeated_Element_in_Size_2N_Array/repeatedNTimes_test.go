package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want int
	}{
		{
			name: "base case",
			got:  []int{1, 2, 3, 3},
			want: 3,
		},
		{
			name: "empty string case",
			got:  []int{2, 1, 2, 5, 3, 2},
			want: 2,
		},
		{
			name: "second case",
			got:  []int{5, 1, 5, 2, 5, 3, 5, 4},
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, repeatedNTimes(test.got))
		})
	}
}
