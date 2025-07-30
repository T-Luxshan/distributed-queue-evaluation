package main

import (
	"fmt"
	"github.com/T-Luxshan/distributed-queue-evaluation/asynq/tasks"
	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func asynqImp(client *asynq.Client) {

	userID := 456
	requestID := "req-456"

	fare := 100

	tasks.EnqueueTask(client, userID, fare, requestID)
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing redis client:", err)
		}
	}(client)
	asynqImp(client)
}
