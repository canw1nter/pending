package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestHelloWorldPublisher(t *testing.T) {
	conn, _ := amqp.Dial("amqp://root:123456@localhost:5672")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch.PublishWithContext(ctx, "", "hello.world", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("What are you doing?"),
	})
}

func TestHelloWorldConsumer(t *testing.T) {
	conn, _ := amqp.Dial("amqp://root:123456@localhost:5672")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	msgs, _ := ch.Consume("hello.world", "", true, false, false, false, nil)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Receive message: %s\n", d.Body)
		}
	}()

	<-forever
}
