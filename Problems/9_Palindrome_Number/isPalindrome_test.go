package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want bool
	}{
		{
			name: "valid",
			got:  1234321,
			want: true,
		},
		{
			name: "invalid",
			got:  101929485,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, isPalindrome(test.got))
		})
	}
}
