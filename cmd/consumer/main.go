package main

import (
	"github.com/vieirinhasantana/kafka-go/pkg/kafka"
)

func main() {
	kafkaProcessor := kafka.NewKafkaProcessor()
	kafkaProcessor.BootstrapServers = "kafka:9092"
	kafkaProcessor.ConsumerTopics = []string{"products"}
	kafkaProcessor.ConsumerGroup = "consumergo"

	kafkaProcessor.Consume()
}
