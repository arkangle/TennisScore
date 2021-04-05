package tennis

import (
	"encoding/json"
)

type State struct {
	Game  []int `json:"game"`
	Deuce int   `json:"deuce"`
	Set   []int `json:"set"`
	Match []int `json:"match"`
}

func NewStateFromJson(str string) (State, error) {
	response := State{}
	err := json.Unmarshal([]byte(str), &response)
	return response, err
}

func (s State) ToJson() (string, error) {
	response, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
