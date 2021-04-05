package tennis

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStateFromJson(t *testing.T) {
	template := `
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
		hasError bool
	}{
		{[]int{0, 0}, 0, []int{0, 0}, []int{0, 0}, false},
		{[]int{15, 30}, 1, []int{5, 4}, []int{5, 4}, false},
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
		)
		state, err := NewStateFromJson(jsonString)
		if err != nil {
			assert.True(t, test.hasError, err)
		}
		assert.Equal(t, test.game, state.Game)
		assert.Equal(t, test.deuce, state.Deuce)
		assert.Equal(t, test.set, state.Set)
		assert.Equal(t, test.match, state.Match)
	}
}

func TestStateToJson(t *testing.T) {
	template := `
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
		hasError bool
	}{
		{[]int{0, 0}, 0, []int{0, 0}, []int{0, 0}, false},
		{[]int{15, 30}, 1, []int{5, 4}, []int{5, 4}, false},
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
		)
		state, _ := NewStateFromJson(jsonString)
		whitespaceReplacer := strings.NewReplacer(" ", "", "\r", "", "\n", "", "\t", "")
		json, err := state.ToJson()
		if err != nil {
			assert.True(t, test.hasError, err)
		}
		assert.Equal(t, whitespaceReplacer.Replace(jsonString), json)

	}
}
