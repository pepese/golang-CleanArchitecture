package app

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AppName     string   `envconfig:"APP_NAME" default:"ca-app"`
	AppVersion  string   `envconfig:"APP_VERSION" default:"undefined"`
	RdbType     string   `envconfig:"RDB_TYPE" default:"mysql"`
	RdbUser     string   `envconfig:"RDB_USER" default:"testuser"`
	RdbPassword string   `envconfig:"RDB_PASSWORD" default:"testpass"`
	RdbProtocol string   `envconfig:"RDB_PROTOCOL" default:"tcp"`
	RdbHost     string   `envconfig:"RDB_HOST" default:"mysql:3306"`
	RdbName     string   `envconfig:"RDB_NAME" default:"testdb"`
	KafkaAddrs  []string `envconfig:"KAFKA_ADDRS" default:"kafka:9092"`
	KafkaTopic  string   `envconfig:"KAFKA_TOPIC" default:"users"`
}

var (
	conf  *config
	cOnce sync.Once
)

func Config() *config {
	initConfig()
	return conf
}

func initConfig() {
	cOnce.Do(func() {
		conf = &config{}
		envconfig.Process("", conf)
	})
}
