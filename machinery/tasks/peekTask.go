package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	MyConfig "github.com/T-Luxshan/distributed-queue-evaluation/machinery/config"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func PeekMessage() (error, time.Duration) {
	conn, ch, err := MyConfig.ConnectToRabbitMQ()
	if err != nil {
		return err, -1
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)
	startPeek := time.Now()
	// Get a message
	msg, ok, err := ch.Get(
		"process_tasks",
		false,
	)
	if err != nil {
		return fmt.Errorf("failed to get message: %w", err), -1
	}

	if !ok {
		return nil, -1
	}
	var task tasks.Signature
	err = json.Unmarshal(msg.Body, &task)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
	} else {
		if len(task.Args) >= 2 {
			fmt.Printf("User ID: %v, Request ID: %v\n", task.Args[0].Value, task.Args[1].Value)
		} else {
			log.Println("Not enough arguments in task")
		}
	}
	// Requeue the message
	err = ch.Nack(msg.DeliveryTag, false, true) //requeue
	if err != nil {
		return errors.New(fmt.Sprintf("failed to requeue message: %v", err)), -1
	}
	elapsedPeek := time.Since(startPeek)
	fmt.Println(elapsedPeek)
	return nil, elapsedPeek
}
