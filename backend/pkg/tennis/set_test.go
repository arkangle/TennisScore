package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetNext(t *testing.T) {
	set := NewStandardSet()
	var tests = []struct {
		current []int
		winner  int
		next    []int
		won     bool
	}{
		{[]int{0, 0}, 0, []int{1, 0}, false},
		{[]int{3, 0}, 0, []int{4, 0}, false},
		{[]int{5, 0}, 0, []int{6, 0}, true},
		{[]int{0, 5}, 1, []int{0, 6}, true},
		{[]int{5, 5}, 1, []int{5, 6}, false},
		{[]int{5, 6}, 0, []int{6, 6}, false},
		{[]int{5, 6}, 1, []int{5, 7}, true},
		{[]int{6, 6}, 1, []int{6, 7}, true},
	}
	for _, test := range tests {
		next, won := set.Next(test.current, test.winner)
		assert.Equal(t, test.next, next, test)
		assert.Equal(t, test.won, won, test)
	}
}
