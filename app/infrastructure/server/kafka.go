package server

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/interface/controller"
)

type kafkaConsumer struct{}

func (kc *kafkaConsumer) Run() {
	conf := app.Config()
	// consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	consumer, err := sarama.NewConsumer(conf.KafkaAddrs, nil)
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
	partitionConsumer, err := consumer.ConsumePartition(conf.KafkaTopic, 0, sarama.OffsetOldest)
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

	router := controller.NewKafkaRouter()

	for {
		select {
		// (3)
		case msg := <-partitionConsumer.Messages():
			log.Printf("topic: %s, offset: %d, key: %s, value: %s, blockTime: %s, Headers: %v, partition: %d\n",
				msg.Topic, msg.Offset, msg.Key, msg.Value, msg.BlockTimestamp, msg.Headers, msg.Partition)
			var input controller.UsersTopic
			if err := json.Unmarshal([]byte(msg.Value), &input); err != nil {
				log.Println("Kafka Users Topic Parse Error.")
				log.Println(err)
			} else {
				router.KafkaRouter(input)
			}
		case <-signals:
			// 終了
			return
		}
	}
}

func NewKafkaConsumer() *kafkaConsumer {
	return &kafkaConsumer{}
}
