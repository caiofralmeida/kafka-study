package main

import (
	"context"

	"github.com/segmentio/kafka-go"
	"fmt"
)

func main() {
	const (
		broker    = "localhost:9092"
		protocol  = "tcp"
		topic     = "jamal"
		partition = 0
	)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())

		if err != nil {
			panic(err)
		}

		fmt.Printf("receiving message [%s: %s]\n", string(m.Key), string(m.Value))
	}
}