package main

import (
	"fmt"
	"kafkaconsumer/kafkaconsumer"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}

	configFile := os.Args[1]
	config := kafkaconsumer.ReadConfig(configFile)
	topic := "purchases"
	timeout := 100
	numOfWorkers := 10
	minCommitCount := 1

	// Define your message processing function
	processMessage := func(message *kafka.Message) {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}

	// Define a stop channel for the main program
	stopChan := make(chan struct{})

	// Run the consumer in a goroutine
	consumer := kafkaconsumer.Consumer{
		Topic:          topic,
		ProcessMessage: processMessage,
		NumOfWorkers:   numOfWorkers,
		Timeout:        timeout,
		MinCommitCount: minCommitCount,
		Config:         config,
		StopChan:       stopChan,
	}
	consumer.Execute()
}
