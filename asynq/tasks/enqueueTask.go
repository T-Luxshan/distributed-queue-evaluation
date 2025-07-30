package tasks

import (
	"errors"
	"github.com/hibiken/asynq"
	"log"
)

const RedisAddr = "127.0.0.1:6379"

func chooseQueue(fair int) (string, error) {
	if fair <= 0 {
		return "", errors.New("fair must be greater than zero")
	}
	switch {
	case fair >= 1500:
		return "critical", nil
	case fair >= 500:
		return "default", nil
	default:
		return "low", nil
	}
}

func EnqueueTask(client *asynq.Client, userID, fare int, rideID string) {

	queueName, err := chooseQueue(fare)
	if err != nil {
		log.Fatalf("Failed to choose queue: %v", err)
	}

	task, err := NewUserRequestTask(userID, rideID)
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	// Enqueue task with queue name
	info, err := client.Enqueue(task, asynq.Queue(queueName))
	if err != nil {
		log.Printf("Failed to enqueue task: %v", err)
	}

	log.Printf("Task enqueued: id=%s queue=%s", info.ID, info.Queue)
}
