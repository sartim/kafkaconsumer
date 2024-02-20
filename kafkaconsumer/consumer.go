package kafkaconsumer

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// type for message processing function
type MessageProcessor func(message *kafka.Message)

type Consumer struct {
	Topic          string
	ProcessMessage MessageProcessor
	NumOfWorkers   int
	Timeout        int
	MinCommitCount int
	Config         kafka.ConfigMap
	StopChan       <-chan struct{}
}

// Consumer consumes messages from Kafka and processes
// them using the provided message processing function.
func (c Consumer) Execute() {

	consumer, err := kafka.NewConsumer(&c.Config)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	err = consumer.SubscribeTopics([]string{c.Topic}, nil)

	// Create a buffered channel to limit the number of messages processed simultaneously
	workerPool := make(chan struct{}, c.NumOfWorkers)

	// Process messages
	msg_count := 0
	for {
		select {
		case <-c.StopChan:
			// If stop signal received, close the consumer and return
			consumer.Close()
			return
		default:
			ev := consumer.Poll(c.Timeout)
			switch e := ev.(type) {
			case *kafka.Message:
				msg_count++
				if msg_count%c.MinCommitCount == 0 {
					go func() {
						offsets, err := consumer.Commit()
						fmt.Println(offsets, err)
					}()
				}

				// Pass the message to a worker goroutine for processing
				workerPool <- struct{}{}
				go func(message *kafka.Message) {
					defer func() {
						<-workerPool // Release the worker slot
					}()
					// Application-specific message processing
					c.ProcessMessage(message)
				}(e)
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				return // Return if error encountered
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}
}
