package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		gotA string
		gotB string
		want string
	}{
		{
			name: "first case",
			gotA: "11",
			gotB: "1",
			want: "100",
		},
		{
			name: "second case",
			gotA: "1010",
			gotB: "1011",
			want: "10101",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, addBinary(test.gotA, test.gotB))
		})
	}
}
