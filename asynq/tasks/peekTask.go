package tasks

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
)

// PeekNextTask fetches the next task in the queue without removing it
func PeekNextTask(redisAddr, queueName string) error {
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer rdb.Close()

	// Asynq stores pending tasks in a **list**, so we use LIndex
	queueKey := fmt.Sprintf("asynq:{%s}:pending", queueName)

	// Get the first task from the list (index 0)
	taskID, err := rdb.LIndex(ctx, queueKey, 0).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("No tasks available in the %s queue.\n", queueName)
		} else {
			log.Printf("Error fetching task from queue: %v", err)
		}
		return err
	}

	// Fetch task details using HGETALL
	taskKey := fmt.Sprintf("asynq:{%s}:t:%s", queueName, taskID)
	taskData, err := rdb.HGetAll(ctx, taskKey).Result()
	if err != nil {
		log.Printf("Error fetching task data: %v", err)
		return err
	}

	if len(taskData) == 0 {
		log.Println("Task details not found in Redis.")
		return nil
	}

	// Log the task data (mainly the `msg` field, which contains the payload)
	log.Printf("Next task in queue: %v", strings.TrimSpace(taskData["msg"]))

	return nil
}
