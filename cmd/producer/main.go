package main

import "github.com/vieirinhasantana/kafka-go/pkg/kafka"

func main() {
	kafkaProcessor := kafka.NewKafkaProcessor()
	kafkaProcessor.BootstrapServers = "kafka:9092"
	kafkaProcessor.Topic = "products"

	kafkaProcessor.Producer()
}
