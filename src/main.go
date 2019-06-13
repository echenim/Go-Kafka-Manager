package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	consumer "github.com/echenim/kafkamanager/src/manager"
)

func main() {

	//initialize log for sarama
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	cg, slipup := consumer.InitConsumer()
	if slipup != nil {
		fmt.Println("error occured in consumer group: ", slipup.Error())
		os.Exit(1)
	}
	defer cg.Close()
	consumer.Consume(cg)
}
