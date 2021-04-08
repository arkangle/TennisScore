package tennis

import (
	"errors"
	"math"
)

// TODO: no add scroring
// TODO: tiebreaker 7, but by 2
// TODO: tiebreaker 10 instead of 3rd set
type Game struct {
	Sequence []int
	WinByTwo bool
}

func NewStandardGame() Game {
	return Game{
		[]int{0, 15, 30, 40},
		true,
	}
}

func (g Game) Next(current []int, deuce int, winner int) ([]int, int, bool, error) {
	won := false
	indexes, err := g.Indexes(current)
	if err != nil {
		return []int{}, 0, won, err
	}
	indexes[winner] += 1
	if indexes[winner] == len(g.Sequence) {
		deuce, won = g.Deuce(indexes, deuce)
		if won {
			return []int{0, 0}, deuce, won, nil
		}
		return current, deuce, won, nil
	}
	next, err := g.Values(indexes)
	return next, 0, won, err
}

func (g Game) Deuce(indexes []int, deuce int) (int, bool) {
	diff := indexes[0] - indexes[1]
	if g.WinByTwo && math.Abs(float64(diff)) == 1 {
		deuce += diff
		if math.Abs(float64(deuce)) > 1 {
			return deuce, true
		}
		return deuce, false
	}
	return 0, true
}

func (g Game) Index(value int) (int, error) {
	for i, seq := range g.Sequence {
		if seq == value {
			return i, nil
		}
	}
	return -1, errors.New("invalid sequence")
}

func (g Game) Indexes(values []int) ([]int, error) {
	indexes := []int{}
	for _, value := range values {
		index, err := g.Index(value)
		if err != nil {
			return []int{}, err
		}
		indexes = append(indexes, index)
	}
	return indexes, nil
}

func (g Game) Values(indexes []int) ([]int, error) {
	values := []int{}
	count := len(g.Sequence)
	for _, index := range indexes {
		if index >= count {
			return values, errors.New("invalid index")
		}
		value := g.Sequence[index]
		values = append(values, value)
	}
	return values, nil
}
