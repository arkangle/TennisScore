package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchNext(t *testing.T) {
	match := NewStandardMatch()
	var tests = []struct {
		current []int
		winner  int
		next    []int
		won     bool
	}{
		{[]int{0, 0}, 0, []int{1, 0}, false},
		{[]int{1, 0}, 0, []int{2, 0}, true},
		{[]int{0, 0}, 1, []int{0, 1}, false},
		{[]int{0, 1}, 1, []int{0, 2}, true},
		{[]int{1, 1}, 1, []int{1, 2}, true},
	}
	for _, test := range tests {
		next, won := match.Next(test.current, test.winner)
		assert.Equal(t, test.next, next, test)
		assert.Equal(t, test.won, won, test)
	}
}
