package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsNearbyDuplicatex(t *testing.T) {
	tests := []struct {
		name    string
		gotNums []int
		gotK    int
		want    bool
	}{
		{
			name:    "first case",
			gotNums: []int{1, 2, 3, 1},
			gotK:    3,
			want:    true,
		},
		{
			name:    "second case",
			gotNums: []int{1, 0, 1, 1},
			gotK:    1,
			want:    true,
		},
		{
			name:    "third case",
			gotNums: []int{1, 2, 3, 1, 2, 3},
			gotK:    2,
			want:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, containsNearbyDuplicate(test.gotNums, test.gotK))
		})
	}
}
