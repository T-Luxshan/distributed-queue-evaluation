package tasks

import (
	"github.com/hibiken/asynq"
	"log"
	"math/rand"
)

//const RedisAddr = "127.0.0.1:6379"

func chooseQueue() string {

	queueTypes := []string{"critical", "default", "low"}
	randomQ := queueTypes[rand.Intn(len(queueTypes))]

	return randomQ
}

func EnqueueTask(client *asynq.Client, userID int, rideID string) {

	queueName := chooseQueue()

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
