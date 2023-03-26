package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name          string
		gotRansomNote string
		gotMagazine   string
		want          bool
	}{
		{
			name:          "first case",
			gotRansomNote: "a",
			gotMagazine:   "b",
			want:          false,
		},
		{
			name:          "first case",
			gotRansomNote: "aa",
			gotMagazine:   "ab",
			want:          false,
		},
		{
			name:          "third case",
			gotRansomNote: "aa",
			gotMagazine:   "aab",
			want:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, canConstruct(test.gotRansomNote, test.gotMagazine))
		})
	}
}
