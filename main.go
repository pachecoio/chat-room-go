package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func main() {
	topic := "chat-room"
	brokers := []string{"localhost:19092"}
	admin, err := NewAdmin(brokers)
	if err != nil {
		panic(err)
	}
	defer admin.Close()

	exists, err := admin.TopicExists(topic)
	if err != nil {
		panic(err)
	}

	if !exists {
		err = admin.CreateTopic(topic, 1, 1)
		if err != nil {
			panic(err)
		}
	}

	// Set user name

	fmt.Print("Enter your name: ")
	username := ""
	fmt.Scanln(&username)

	producer, err := NewProducer(brokers, topic)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	consumer := NewConsumer(brokers, topic)
	defer consumer.Close()

	go consumer.PrintMessages()

	fmt.Println("Connected. Press Ctrl+C to exit.")

	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		producer.Send(username, message)
	}

}
