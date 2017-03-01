package rabbit

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func Republish(deadQueue, exchange string) {
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer channel.Close()

	for {
		msg, ok, err := channel.Get(deadQueue, false)
		if err != nil {
			log.Fatal(err.Error())
		}

		if !ok {
			log.Println("No more messages")
			break
		}

		headers := msg.Headers
		headers["x-retry-count"] = "0"

		pub := amqp.Publishing{
			Headers:         headers,
			ContentType:     msg.ContentType,
			ContentEncoding: msg.ContentEncoding,
			DeliveryMode:    msg.DeliveryMode,
			Priority:        msg.Priority,
			Expiration:      msg.Expiration,
			Timestamp:       time.Now(),
			Body:            msg.Body,
		}

		err = channel.Publish(exchange, msg.RoutingKey, false, false, pub)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Processed Message")
		msg.Ack(false)
	}
}
