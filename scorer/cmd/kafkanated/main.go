package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddress := "kafka:9092"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "score",
	})
	/*
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{brokerAddress},
			Topic:   "websocket",
		})
	*/
	ctx := context.Background()
	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}
		fmt.Printf("Message[%d] %s = %s\n", message.Offset, message.Key, message.Value)
	}

}
