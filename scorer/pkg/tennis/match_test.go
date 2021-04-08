package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchNext(t *testing.T) {
	var tests = []struct {
		sets     int
		current  []int
		winner   int
		next     []int
		won      bool
		hasError bool
	}{
		{3, []int{0, 0}, 0, []int{1, 0}, false, false},
		{3, []int{1, 0}, 0, []int{2, 0}, true, false},
		{3, []int{0, 0}, 1, []int{0, 1}, false, false},
		{3, []int{0, 1}, 1, []int{0, 2}, true, false},
		{3, []int{1, 1}, 1, []int{1, 2}, true, false},
		{3, []int{1, 2}, 1, []int{0, 0}, false, true},
		{3, []int{3, 2}, 1, []int{0, 0}, false, true},
		{5, []int{2, 2}, 1, []int{2, 3}, true, false},
		{5, []int{4, 2}, 1, []int{0, 0}, false, true},
		{5, []int{3, 2}, 1, []int{0, 0}, false, true},
	}
	for _, test := range tests {
		match := Match{test.sets}
		next, won, err := match.Next(test.current, test.winner)
		if err != nil {
			assert.True(t, test.hasError)
		}
		assert.Equal(t, test.next, next, test)
		assert.Equal(t, test.won, won, test)
	}
}

func TestMatchBestOf(t *testing.T) {
	var tests = []struct {
		sets   int
		bestOf int
	}{
		{3, 2},
		{5, 3},
		{7, 4},
	}
	for _, test := range tests {
		match := Match{test.sets}
		bestOf := match.BestOf()
		assert.Equal(t, test.bestOf, bestOf)
	}
}
