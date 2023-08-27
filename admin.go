package main

import (
	"context"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Admin struct {
	client *kadm.Client
}

func NewAdmin(brokers []string) (*Admin, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, err
	}
	admin := kadm.NewClient(client)
	return &Admin{client: admin}, nil
}

func (a *Admin) TopicExists(topic string) (bool, error) {
	ctx := context.Background()
	topicsMetadata, err := a.client.ListTopics(ctx)
	if err != nil {
		return false, err
	}
	for _, topicMetadata := range topicsMetadata {
		if topicMetadata.Topic == topic {
			return true, nil
		}
	}
	return false, nil
}

func (a *Admin) CreateTopic(topic string, partitions int, replicationFactor int) error {
	ctx := context.Background()
	resp, err := a.client.CreateTopics(ctx, 1, 1, nil, topic)
	if err != nil {
		return err
	}
	for _, ctr := range resp {
		if ctr.Err != nil {
			return ctr.Err
		}
	}
	return nil
}

func (a *Admin) Close() {
	a.client.Close()
}
