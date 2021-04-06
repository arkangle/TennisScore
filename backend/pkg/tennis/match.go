package tennis

import (
	"errors"
)

type Match struct {
	Sets int
}

func NewStandardMatch() Match {
	return Match{
		Sets: 3,
	}
}

func (m Match) Next(current []int, winner int) ([]int, bool, error) {
	bestOf := m.BestOf()
	next := []int{current[0], current[1]}
	if next[0] < 0 || next[1] < 0 || next[0] >= bestOf || next[1] >= bestOf {
		return []int{0, 0}, false, errors.New("invalid match sets")
	}
	next[winner] += 1
	won := next[winner] == bestOf
	return next, won, nil
}

func (m Match) BestOf() int {
	return (int)(m.Sets/2) + 1
}
