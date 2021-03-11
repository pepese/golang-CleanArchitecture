package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type kafkaConsumer struct{}

func (kc *kafkaConsumer) Run() {
	tmp := "default_topic"
	topic := &tmp
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("consumer created")
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("commence consuming")
	partitionConsumer, err := consumer.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		// (3)
		case msg := <-partitionConsumer.Messages():
			log.Printf("topic: %s, offset: %d, key: %s, value: %s, blockTime: %s, Headers: %v, partition: %d\n",
				msg.Topic, msg.Offset, msg.Key, msg.Value, msg.BlockTimestamp, msg.Headers, msg.Partition)
		case <-signals:
			// 終了
			return
		}
	}
}

func NewKafkaConsumer() *kafkaConsumer {
	return &kafkaConsumer{}
}
