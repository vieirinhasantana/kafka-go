package kafka

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
	ConsumerTopics   []string
	BootstrapServers string
	ConsumerGroup    string
	Topic            string
}

func NewKafkaProcessor() *KafkaProcessor {
	return &KafkaProcessor{}
}

func (k *KafkaProcessor) Consume() {
	c, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": k.BootstrapServers,
		"group.id":          k.ConsumerGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(k.ConsumerTopics, nil)
	log.Printf("Kafka consumer has been started on host %s", k.BootstrapServers)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (k *KafkaProcessor) Producer() {
	p, err := ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": k.BootstrapServers,
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&ckafka.Message{
			TopicPartition: ckafka.TopicPartition{Topic: &k.Topic, Partition: ckafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
		log.Printf("New event send. %s", word)
	}

	p.Flush(15 * 1000)

}
