package tennis

type Match struct {
	Sets int
}

func NewStandardMatch() Match {
	return Match{
		Sets: 3,
	}
}
func (m Match) Next(current []int, winner int) ([]int, bool) {
	next := []int{current[0], current[1]}
	next[winner] += 1
	won := next[winner] == (m.Sets - 1)
	return next, won
}
