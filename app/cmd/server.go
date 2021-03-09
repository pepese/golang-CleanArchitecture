// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/pepese/golang-CleanArchitecture/app/infrastructure/server"
	"github.com/spf13/cobra"
)

var t string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().StringVarP(&t, "type", "t", "http", "")
}

func runServer(cmd *cobra.Command, args []string) {
	if t == "http" {
		srv := server.NewHttpServer()
		srv.Run()
	} else if t == "gin" {
		srv := server.NewGinServer()
		srv.Run()
	} else if t == "grpc" {
		srv := server.NewGrpcServer()
		srv.Run()
	} else if t == "kafka" {
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
				log.Printf("topic: %s, offset: %d, key: %s, value: %s\n",
					msg.Topic, msg.Offset, msg.Key, msg.Value)
			case <-signals:
				// 終了
				return
			}
		}
	}
}
