package main

import (
	"kafkaconsumer/kafkaconsumer"
	"testing"
)

// TestReadConfig tests the ReadConfig function.
func TestReadConfig(t *testing.T) {
	configFile := "../config.properties"

	config := kafkaconsumer.ReadConfig(configFile)

	expectedBootstrapServers := "localhost:9092"
	if config["bootstrap.servers"] != expectedBootstrapServers {
		t.Errorf("Expected bootstrap servers: %s, got: %s", expectedBootstrapServers, config["bootstrap.servers"])
	}

	expectedGroupID := "default-consumer-group"
	if config["group.id"] != expectedGroupID {
		t.Errorf("Expected group ID: %s, got: %s", expectedGroupID, config["group.id"])
	}

	expectedAutoOffsetReset := "earliest"
	if config["auto.offset.reset"] != expectedAutoOffsetReset {
		t.Errorf("Expected auto.offset.reset: %s, got: %s", expectedAutoOffsetReset, config["auto.offset.reset"])
	}
}
