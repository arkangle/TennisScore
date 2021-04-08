package tennis

import (
	"encoding/json"
)

type Score struct {
	State  State `json:"state"`
	Winner int   `json:"winner"`
	Id     int   `json:"id"`
}

type State struct {
	Game  []int `json:"game"`
	Deuce int   `json:"deuce"`
	Set   []int `json:"set"`
	Match []int `json:"match"`
}

func NewScoreFromJson(str string) (Score, error) {
	response := Score{}
	err := json.Unmarshal([]byte(str), &response)
	return response, err
}

func (s Score) ToJson() (string, error) {
	response, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
