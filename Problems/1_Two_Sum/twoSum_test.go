package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name      string
		gotNums   []int
		gotTarget int
		want      []int
	}{
		{
			name:      "first case",
			gotNums:   []int{2, 7, 11, 15},
			gotTarget: 9,
			want:      []int{0, 1},
		},
		{
			name:      "second case",
			gotNums:   []int{3, 2, 4},
			gotTarget: 6,
			want:      []int{1, 2},
		},
		{
			name:      "third case",
			gotNums:   []int{3, 3},
			gotTarget: 6,
			want:      []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, twoSum(test.gotNums, test.gotTarget))
		})
	}
}
