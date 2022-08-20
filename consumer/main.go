package main

import (
	"github.com/korasdor/rabbitmq-go/consumer/consts"
	"github.com/korasdor/rabbitmq-go/consumer/handlers"
	"github.com/korasdor/rabbitmq-go/consumer/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")

	exampleQueue := utils.RMQConsumer{
		Queue:            consts.EXAMPLE_QUEUE,
		ConnectionString: connectionString,
		MsgHandler:       handlers.HandleExample,
	}

	forever := make(chan bool)
	go exampleQueue.Consume()

	<-forever
}
