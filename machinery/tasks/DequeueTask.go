package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	MyConfig "github.com/T-Luxshan/distributed-queue-evaluation/machinery/config"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

// DequeueMessage Dequeues a message from the queue.
func DequeueMessage() (error, time.Duration) {
	conn, ch, err := MyConfig.ConnectToRabbitMQ()
	if err != nil {
		return nil, -1
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
	var wg sync.WaitGroup
	startDequeue := time.Now()
	isQueueEmpty := make(chan bool)
	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Get a message, and acknowledge it.
			msg, ok, err := ch.Get(
				"process_tasks",
				false,
			)
			if err != nil {
				_ = fmt.Errorf("failed to get message: %w", err)
				return
			}

			if !ok {
				log.Println("No task found in the queue")
				isQueueEmpty <- true
				return
			} else {
				isQueueEmpty <- false
			}

			var task tasks.Signature
			err = json.Unmarshal(msg.Body, &task)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
			} else {
				if len(task.Args) >= 2 {
					fmt.Printf("Task dequeued, User ID: %v, Request ID: %v\n", task.Args[0].Value, task.Args[1].Value)
				} else {
					log.Println("Not enough arguments in task")
				}
			}
			err = ch.Ack(msg.DeliveryTag, false)
			if err != nil {
				return
			}
		}()
		if <-isQueueEmpty {
			break
		}
	}
	wg.Wait()
	elapsedDequeue := time.Since(startDequeue)
	fmt.Println(elapsedDequeue)

	return nil, elapsedDequeue
}
