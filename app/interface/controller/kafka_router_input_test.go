package controller

import (
	"encoding/json"
	"log"
	"testing"
)

func TestKafkaRouterInput(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		msg := `{"method":"create","message":{"first_name":"first","last_name":"last"}}`
		var input UsersTopic
		if err := json.Unmarshal([]byte(msg), &input); err != nil {
			log.Println("Kafka Users Topic Parse Error.")
			log.Println(err)
		} else {
			log.Printf("input: %v\n", input)
		}
	})
}
