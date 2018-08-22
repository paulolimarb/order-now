package queue

import (
	"log"
	"order-now/order-consumer/config"

	"order-now/order-consumer/repository"

	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ReadMessage() {
	conn, err := amqp.Dial(config.Get().RabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"orders",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	consumerName := uuid.NewV4()
	msgs, err := ch.Consume(
		q.Name,
		consumerName.String(),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			v := repository.CreateOrder(d.Body)

			if v {
				d.Ack(true)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
