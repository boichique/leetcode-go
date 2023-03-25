package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want bool
	}{
		{
			name: "first case",
			got:  "()",
			want: true,
		},
		{
			name: "second case",
			got:  "()[]{}",
			want: true,
		},
		{
			name: "third case",
			got:  "(]",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, isValid(test.got))
		})
	}
}
