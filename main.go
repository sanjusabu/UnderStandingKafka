package main

import (
	"github.com/sanjusabu/UnderStandingKafka/consumer"
	"github.com/sanjusabu/UnderStandingKafka/utils"
)

func main() {
	// add producer and consumer code here
	// creates a new producer instance
	conf := utils.ReadConfig()
	// fmt.Print("check2", conf)

	topic := "topic_0" // This topic should be created in confluent cloud by you before hand

	// go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	// producer.ProduceMessages(&conf, topic)
	consumer.ConsumeMessages(conf, topic)
}
