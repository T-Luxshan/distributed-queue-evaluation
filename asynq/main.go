package main

import (
	"fmt"
	"github.com/T-Luxshan/distributed-queue-evaluation/asynq/tasks"
	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func asynqImp(client *asynq.Client) {

	userID := 123
	requestID := "req-123"

	tasks.EnqueueTask(client, userID, requestID)

	err := tasks.PeekNextTask(redisAddr, "low")
	if err != nil {
		fmt.Println(err)
	}
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
