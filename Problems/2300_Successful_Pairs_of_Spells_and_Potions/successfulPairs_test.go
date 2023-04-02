package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessfulPairs(t *testing.T) {
	tests := []struct {
		name       string
		gotSpells  []int
		gotPotions []int
		gotSuccess int64
		want       []int
	}{
		{
			name:       "first case",
			gotSpells:  []int{5, 1, 3},
			gotPotions: []int{1, 2, 3, 4, 5},
			gotSuccess: 7,
			want:       []int{4, 0, 3},
		},
		{
			name:       "second case",
			gotSpells:  []int{3, 1, 2},
			gotPotions: []int{8, 5, 8},
			gotSuccess: 16,
			want:       []int{2, 0, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, successfulPairs(test.gotSpells, test.gotPotions, test.gotSuccess))
		})
	}
}
