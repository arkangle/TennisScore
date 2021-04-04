package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameIndex(t *testing.T) {
	game := NewStandardGame()
	var tests = []struct {
		value    int
		index    int
		hasError bool
	}{
		{0, 0, false},
		{15, 1, false},
		{30, 2, false},
		{40, 3, false},
		{5, -1, true},
	}
	for _, test := range tests {
		index, err := game.Index(test.value)
		if err != nil {
			assert.True(t, test.hasError)
		}
		assert.Equal(t, test.index, index)
	}
}

func TestGameDeuce(t *testing.T) {
	game := NewStandardGame()
	var tests = []struct {
		indexes   []int
		deuce     int
		deuceNext int
		won       bool
	}{
		{[]int{0, 1}, 0, -1, false},
		{[]int{1, 0}, 0, 1, false},
		{[]int{1, 0}, 1, 2, true},
		{[]int{1, 0}, 1, 2, true},
	}
	for _, test := range tests {
		deuceNext, won := game.Deuce(test.indexes, test.deuce)
		assert.Equal(t, test.deuceNext, deuceNext)
		assert.Equal(t, test.won, won)
	}
}

func TestGameIndexes(t *testing.T) {
	game := NewStandardGame()
	var tests = []struct {
		values   []int
		indexes  []int
		hasError bool
	}{
		{[]int{0, 0}, []int{0, 0}, false},
		{[]int{15, 15}, []int{1, 1}, false},
		{[]int{30, 30}, []int{2, 2}, false},
		{[]int{40, 40}, []int{3, 3}, false},
		{[]int{30, 40}, []int{2, 3}, false},
		{[]int{5, 30}, []int{-1, 3}, true},
		{[]int{15, 5}, []int{-1, 3}, true},
	}
	for _, test := range tests {
		indexes, err := game.Indexes(test.values)
		if err != nil {
			assert.True(t, test.hasError)
		} else {
			assert.Equal(t, test.indexes, indexes)
		}
	}
}

func TestGameValues(t *testing.T) {
	game := NewStandardGame()
	var tests = []struct {
		indexes  []int
		values   []int
		hasError bool
	}{
		{[]int{0, 0}, []int{0, 0}, false},
		{[]int{1, 1}, []int{15, 15}, false},
		{[]int{2, 2}, []int{30, 30}, false},
		{[]int{3, 3}, []int{40, 40}, false},
		{[]int{2, 3}, []int{30, 40}, false},
		{[]int{5, 2}, []int{}, true},
		{[]int{1, 5}, []int{}, true},
	}
	for _, test := range tests {
		values, err := game.Values(test.indexes)
		if err != nil {
			assert.True(t, test.hasError)
		} else {
			assert.Equal(t, test.values, values)
		}
	}
}

func TestStandardGameNext(t *testing.T) {
	game := NewStandardGame()
	var tests = []struct {
		current   []int
		deuce     int
		next      []int
		winner    int
		deuceNext int
		end       bool
		hasError  bool
	}{
		{[]int{15, 5}, 0, []int{0, 0}, 0, 0, false, true},
		{[]int{0, 0}, 0, []int{15, 0}, 0, 0, false, false},
		{[]int{15, 0}, 0, []int{30, 0}, 0, 0, false, false},
		{[]int{15, 30}, 0, []int{15, 40}, 1, 0, false, false},
		{[]int{40, 30}, 0, []int{0, 0}, 0, 0, true, false},
		{[]int{40, 40}, 0, []int{40, 40}, 0, 1, false, false},
		{[]int{40, 40}, 1, []int{0, 0}, 0, 2, true, false},
		{[]int{40, 40}, 1, []int{40, 40}, 1, 0, false, false},
		{[]int{40, 40}, 0, []int{40, 40}, 1, -1, false, false},
		{[]int{40, 40}, -1, []int{0, 0}, 1, -2, true, false},
	}
	for _, test := range tests {
		next, deuceNext, end, err := game.Next(test.current, test.deuce, test.winner)
		if err != nil {
			assert.True(t, test.hasError)
		} else {
			assert.Equal(t, test.next, next)
			assert.Equal(t, test.end, end)
			assert.Equal(t, test.deuceNext, deuceNext)
		}
	}
}
