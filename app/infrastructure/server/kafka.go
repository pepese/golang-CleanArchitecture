package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/interface/controller"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
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

	router := controller.NewKafkaRouter(nil)

	for {
		select {
		// (3)
		case msg := <-partitionConsumer.Messages():
			// context.Context
			c := context.Background()
			ctx := context.Background()

			// Request ID
			reqId := uuid.NewV4().String()

			// Logger
			logger := app.LoggerWithKeyValue("reqId", reqId)
			ctx = app.SetLoggerToContext(ctx, logger)
			c = context.WithValue(c, "ctx", ctx)

			// Access Log
			start := time.Now()
			// アクセスログ（リクエスト）出力
			logger.Infow("===>request",
				fmt.Sprintf("topic: %s, offset: %d, key: %s, value: %s, blockTime: %s, Headers: %v, partition: %d. ",
					msg.Topic, msg.Offset, msg.Key, msg.Value, msg.BlockTimestamp, msg.Headers, msg.Partition),
				"reqId", reqId)

			var input controller.UsersTopic
			if err := json.Unmarshal([]byte(msg.Value), &input); err != nil {
				logger.Infow("Kafka Users Topic Parse Error: ", err)
			} else {
				reply, err := router.KafkaRouter(c, input)

				// アクセスログ（レスポンス）出力
				responseTime := time.Now().Sub(start)
				if err != nil {
					logger.Infow("<==response", zap.Reflect("response", reply), "reqId", reqId, "responseTime", responseTime, "error", err)
				} else {
					logger.Infow("<==response", "reqId", reqId, "responseTime", responseTime)
				}
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
