package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "first case",
			got:  "G()(al)",
			want: "Goal",
		},
		{
			name: "second case",
			got:  "G()()()()(al)",
			want: "Gooooal",
		},
		{
			name: "third case",
			got:  "(al)G(al)()()G",
			want: "alGalooG",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, interpret(test.got))
		})
	}
}
