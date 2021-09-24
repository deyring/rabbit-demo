package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notifications", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	for {
		fmt.Println("Enter your message: ")
		input, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		failOnError(err, "Failed to read input")

		notification := Notification{Message: string(input)}
		body, err := json.Marshal(&notification)
		failOnError(err, "Failed to marshal a task to json")
		err = ch.Publish(
			"notifications", // exchange
			"notifier",      // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         body,
			})
		failOnError(err, "Failed to publish a message")
	}

}

type Notification struct {
	Message string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
