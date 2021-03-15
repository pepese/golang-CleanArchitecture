package controller

import "log"

type kafkaRouter struct{}

func (kr *kafkaRouter) KafkaRouter(in UsersTopic) {
	log.Println("kafkaRouter exec.")
	log.Printf("input: %v\n", in)
	return
}

func NewKafkaRouter() *kafkaRouter {
	return &kafkaRouter{}
}
