package tennis

import (
	"errors"
	"math"
)

type Set struct {
	WinByTwo bool
	Games    int
	Limit    int
}

func NewStandardSet() Set {
	return Set{
		WinByTwo: true,
		Games:    6,
		Limit:    7,
	}
}

func (s Set) Next(current []int, winner int) ([]int, bool, error) {
	next := []int{current[0], current[1]}
	next[winner] += 1
	if next[0] < 0 || next[1] < 0 || next[0] > s.Limit || next[1] > s.Limit {
		return []int{0, 0}, false, errors.New("invalid set")
	}
	diff := math.Abs(float64(next[0] - next[1]))
	if next[winner] >= s.Games && diff > 0 {
		if s.WinByTwo && diff == 1 && next[winner] < s.Limit {
			return next, false, nil
		}
		return next, true, nil
	}
	return next, false, nil
}
