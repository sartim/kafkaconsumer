package main

import (
	"kafkaconsumer/kafkaconsumer"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// TestConsumer tests the Consumer function.// TestConsumer tests the Consumer function.// TestConsumer tests the Consumer function.
func TestConsumer(t *testing.T) {
	configFile := "../config.properties"
	config := kafkaconsumer.ReadConfig(configFile)

	topic := "purchases"
	numOfWorkers := 5
	timeout := 100
	minCommitCount := 1

	processMessage := func(message *kafka.Message) {}

	// Define a stop channel for the test
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
	go func() {
		consumer.Execute()
	}()

	// close the stop channel to stop the consumer
	close(stopChan)
}
