package main

import (
	"context"
	"os"

	"github.com/segmentio/kafka-go"

	"fmt"
)

type msg struct {
	key   []byte
	value []byte
}

func newMsg(k string, v string) msg {
	return msg{[]byte(k), []byte(v)}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("You should give the message key and value to command.")
		fmt.Println("go run producer <key> \"<message>\"")
		os.Exit(1)
	}

	msg := newMsg(os.Args[1], os.Args[2])

	const (
		broker    = "localhost:9092"
		protocol  = "tcp"
		topic     = "jamal"
		partition = 0
	)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	defer w.Close()

	err := w.WriteMessages(context.Background(), kafka.Message{
		Partition: partition,
		Key:       msg.key,
		Value:     msg.value,
	})

	if err != nil {
		panic(err)
	}
}
