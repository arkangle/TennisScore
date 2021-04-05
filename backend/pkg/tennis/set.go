package tennis

import "math"

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

func (s Set) Next(current []int, winner int) ([]int, bool) {
	next := []int{current[0], current[1]}
	next[winner] += 1
	diff := math.Abs(float64(next[0] - next[1]))
	if next[winner] >= s.Games && diff > 0 {
		if s.WinByTwo && diff == 1 && next[winner] < s.Limit {
			return next, false
		}
		return next, true
	}
	return next, false
}
