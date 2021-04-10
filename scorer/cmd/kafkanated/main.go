package main

import (
	"context"
	"fmt"

	"github.com/arkangle/tennisscore/scorer/pkg/tennis"
	"github.com/segmentio/kafka-go"
)

func processMessage(message kafka.Message) tennis.State {
	fmt.Printf("Message[%d] %s = %s\n", message.Offset, message.Key, message.Value)
	msg := string(message.Value)
	score, err := tennis.NewScoreFromJson(msg)
	if err != nil {
		fmt.Printf("Message[%d] %s ERROR: %s\n", message.Offset, message.Key, err)
	}
	game := tennis.NewStandardGame()
	nextGame, nextDeuce, wonGame, _ := game.Next(score.State.Game, score.State.Deuce, score.Winner)
	score.State.Game = nextGame
	score.State.Deuce = nextDeuce
	if wonGame {
		set := tennis.NewStandardSet()
		nextSet, wonSet, _ := set.Next(score.State.Set, score.Winner)
		score.State.Set = nextSet
		if wonSet {
			match := tennis.NewStandardMatch()
			nextMatch, _, _ := match.Next(score.State.Match, score.Winner)
			score.State.Match = nextMatch
		}
	}

	fmt.Printf("Message[%d] %s = %s\n", message.Offset, message.Key, message.Value)
	return score.State
}

func main() {
	brokerAddress := "kafka:9093"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "score",
		GroupID: "scorer",
	})
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{brokerAddress},
		Topic:     "websocket",
		BatchSize: 1,
	})
	ctx := context.Background()
	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}
		nextState := processMessage(message)
		wsMessage, _ := nextState.ToJson()
		writer.WriteMessages(ctx, kafka.Message{
			Key:   message.Key,
			Value: []byte(wsMessage),
		})
		fmt.Printf("Sent To Websocket: %s\n", wsMessage)
	}

}
