package tennis

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScoreFromJson(t *testing.T) {
	template := `
	{
		"state": {
			"game":[%d,%d],
			"deuce": %d,
			"set": [%d,%d],
			"match": [%d,%d]
		},
		"winner": %d,
		"id": %d
	}
	`
	var tests = []struct {
		game     []int
		deuce    int
		set      []int
		match    []int
		winner   int
		id       int
		hasError bool
	}{
		{[]int{0, 0}, 0, []int{0, 0}, []int{0, 0}, 0, 0, false},
		{[]int{15, 30}, 1, []int{5, 4}, []int{5, 4}, 1, 25, false},
	}
	for _, test := range tests {
		jsonString := fmt.Sprintf(template,
			test.game[0],
			test.game[1],
			test.deuce,
			test.set[0],
			test.set[1],
			test.match[0],
			test.match[1],
			test.winner,
			test.id,
		)
		score, err := NewScoreFromJson(jsonString)
		if err != nil {
			assert.True(t, test.hasError, err)
		}
		state := score.State
		assert.Equal(t, test.game, state.Game)
		assert.Equal(t, test.deuce, state.Deuce)
		assert.Equal(t, test.set, state.Set)
		assert.Equal(t, test.match, state.Match)
	}
}

func TestStateToJson(t *testing.T) {
	scoreTemplate := `
	{
		"state": {
			"game":[%d,%d],
			"deuce": %d,
			"set": [%d,%d],
			"match": [%d,%d]
		},
		"winner": %d,
		"id": %d
	}
	`
	stateTemplate := `
	{
		"game":[%d,%d],
		"deuce": %d,
		"set": [%d,%d],
		"match": [%d,%d]
	}
	`
	var tests = []struct {
		game     []int
		deuce    int
		set      []int
		match    []int
		winner   int
		id       int
		hasError bool
	}{
		{[]int{0, 0}, 0, []int{0, 0}, []int{0, 0}, 0, 0, false},
		{[]int{15, 30}, 1, []int{5, 4}, []int{5, 4}, 1, 2, false},
	}
	for _, test := range tests {
		jsonScore := fmt.Sprintf(scoreTemplate,
			test.game[0],
			test.game[1],
			test.deuce,
			test.set[0],
			test.set[1],
			test.match[0],
			test.match[1],
			test.winner,
			test.id,
		)
		score, _ := NewScoreFromJson(jsonScore)
		jsonState := fmt.Sprintf(stateTemplate,
			test.game[0],
			test.game[1],
			test.deuce,
			test.set[0],
			test.set[1],
			test.match[0],
			test.match[1],
		)
		whitespaceReplacer := strings.NewReplacer(" ", "", "\r", "", "\n", "", "\t", "")
		json, err := score.State.ToJson()
		if err != nil {
			assert.True(t, test.hasError, err)
		}
		assert.Equal(t, whitespaceReplacer.Replace(jsonState), json)

	}
}
