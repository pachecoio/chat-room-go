package main

import (
	"context"
	"encoding/json"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
	topic  string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, err
	}
	return &Producer{client: client, topic: topic}, nil
}

func (p *Producer) Send(user, message string) error {
	ctx := context.Background()
	msg := Message{
		User:    user,
		Message: message,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	record := kgo.Record{
		Topic: p.topic,
		Value: msgBytes,
	}
	p.client.Produce(ctx, &record, nil)
	return nil
}

func (p *Producer) Close() {
	p.client.Close()
}
