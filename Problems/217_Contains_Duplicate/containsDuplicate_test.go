package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		got  []int
		want bool
	}{
		{
			name: "first case",
			got:  []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "second case",
			got:  []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "third case",
			got:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, containsDuplicate(test.got))
		})
	}
}
