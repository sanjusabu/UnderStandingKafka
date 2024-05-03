package producer

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceMessages(conf *kafka.ConfigMap, topic string) {
	p, _ := kafka.NewProducer(conf)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	// produces a sample message to the user-created topic
	stoptime := time.Now().Add(90000)
	nums := 1
	for time.Now().Before(stoptime) {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(fmt.Sprintf(`Kafka Message no. : %d`, nums)),
			Value:          []byte(fmt.Sprintf(`Msg No. %d Received by Sanju`, nums)),
		}, nil)
		nums++
	}
	// se9nd any outstanding or buffered messages to the Kafka broker and close the connection
	p.Flush(15 * 1000)
	p.Close()
}
