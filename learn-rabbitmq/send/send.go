package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "fail on connecting to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "fail on open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare channel")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	body := "hello world"
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "failed to publish a message")
	log.Printf("[x] Sent %s\n", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
