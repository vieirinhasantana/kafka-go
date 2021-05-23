package kafka

"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type ConfigureKafka struct {
	servers string
	group   string
	topics  []string
}

func NewConfig(cf ConfigureKafka) *ConfigureKafka {
	return &ConfigureKafka{}
}


func ConnectKafka() {
	c, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": ConfigureKafka.servers,
		"group.id":          ConfigureKafka.group
	})
}
