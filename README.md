# Kafka Consumer Library for Go

[![Language](https://img.shields.io/badge/language-go-blue.svg)](https://github.com/sartim/kafka-consumer)
[![Build Status](https://github.com/sartim/kafka-consumer/workflows/build/badge.svg)](https://github.com/sartim/kafka-consumer)

This Go library provides functionality for consuming messages from Kafka with concurrent message processing.

## Features

- Consumes messages from Kafka topics.
- Supports concurrent message processing using worker pool.
- Simple to use and integrate into your Go applications.

## Installation

To use this library in your Go project, you can simply run:

```sh
go get -u github.com/sartim/kafkaconsumer
```

## Usage

To use this package, import it into your Go project:


```go
import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/sartim/kafkaconsumer"
)
```

## Consumer Function
The Consumer function is used to consume messages from Kafka and process them using the provided message processing function.

```go
type Consumer struct {
	Topic          string
	ProcessMessage MessageProcessor
	NumOfWorkers   int
	Timeout        int
	MinCommitCount int
	Config         kafka.ConfigMap
	StopChan       <-chan struct{}
}

func (c Consumer) Execute()

```

* `Topic`: The Kafka topic to consume messages from.
* `ProcessMessage`: A function that defines how to process the consumed messages.
* `NumOfWorkers`: The number of worker goroutines to limit concurrent message processing.
* `Timeout`: The timeout in milliseconds for polling Kafka messages.
* `minCommitCount`: Commit triggered every `minCommitCount`.
* `Config`: Kafka configuration parameters.
* `StopChan`: Stop channel for the main program.


## Example 

```go
func main() {
    topic := "my_topic"
    consumerGroupID := "my_consumer_group"
    numWorkers := 10
    timeout := 100 

    config := kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        // Add any other Kafka configuration parameters here
    }

	processMessage := func(message *kafka.Message) {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}

	stopChan := make(chan struct{})

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
```

Running Consumer as goroutine

```go
func main() {
    topic := "my_topic"
    numOfWorkers := 10
    timeout := 100 
	minCommitCount := 1
    stopChan := make(chan struct{})

    config := kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        // Add any other Kafka configuration parameters here
    }

	processMessage := func(message *kafka.Message) {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}

	consumer := kafkaconsumer.Consumer{
		Topic:          topic,
		ProcessMessage: processMessage,
		NumOfWorkers:   numOfWorkers,
		Timeout:        timeout,
		MinCommitCount: minCommitCount,
		Config:         config,
		StopChan:       stopChan,
	}
	go consumer.Execute()

    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
    <-sigchan
}
```

## Contribution

Contributions to this package are welcome. Feel free to submit issues or pull requests on the GitHub repository.
