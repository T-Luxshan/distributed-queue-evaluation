package config

import (
	"fmt"
	"github.com/streadway/amqp"
)

// RabbitMQ connection details
const (
	rabbitmqURL = "amqp://guest:guest@localhost:5672/"
)

func ConnectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return conn, ch, nil
}
