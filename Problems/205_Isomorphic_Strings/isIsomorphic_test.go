package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name string
		gotS string
		gotT string
		want bool
	}{
		{
			name: "first case",
			gotS: "egg",
			gotT: "add",
			want: true,
		},
		{
			name: "second case",
			gotS: "foo",
			gotT: "bar",
			want: false,
		},
		{
			name: "third case",
			gotS: "paper",
			gotT: "title",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, isIsomorphic(test.gotS, test.gotT))
		})
	}
}
