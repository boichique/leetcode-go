package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want bool
	}{
		{
			name: "first case",
			got:  28,
			want: true,
		},
		{
			name: "second case",
			got:  7,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, checkPerfectNumber(test.got))
		})
	}
}
