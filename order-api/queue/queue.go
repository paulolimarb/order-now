package queue

import (
	"fmt"
	"log"

	"order-now/order-api/config"

	"github.com/streadway/amqp"
)

func WriteMessage(body string) error {

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
	failOnError(err, "Failed to create queue")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	return err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		fmt.Sprintf("%s: %s", msg, err)
	}
}
