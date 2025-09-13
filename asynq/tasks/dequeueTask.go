package tasks

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
)

// DequeueTask fetches and removes the next task from the queue
func DequeueTask(redisAddr, queueName string) error {
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer rdb.Close()

	// Asynq stores tasks in a list, so we use LPop to dequeue
	queueKey := fmt.Sprintf("asynq:{%s}:pending", queueName)

	// Remove and get the first task from the queue
	taskID, err := rdb.LPop(ctx, queueKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("No tasks available to dequeue.")
		} else {
			log.Printf("Error dequeuing task: %v", err)
		}
		return err
	}

	// Fetch task details
	taskKey := fmt.Sprintf("asynq:{%s}:t:%s", queueName, taskID)
	taskData, err := rdb.HGetAll(ctx, taskKey).Result()
	if err != nil {
		log.Printf("Error fetching dequeued task data: %v", err)
		return err
	}

	if len(taskData) == 0 {
		log.Printf("Dequeued task ID %s not found in Redis.", taskID)
		return nil
	}

	// Extract and clean up the task message
	rawMsg := strings.TrimSpace(taskData["msg"])
	log.Printf("Dequeued task: %s", rawMsg)

	// Delete the task metadata from Redis
	err = rdb.Del(ctx, taskKey).Err()
	if err != nil {
		log.Printf("Error deleting task metadata: %v", err)
		return err
	}

	log.Printf("Task %s successfully dequeued and removed.", taskID)
	return nil
}

func DequeueAll(redisAddr, queueName string) error {
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer rdb.Close()

	// Asynq stores pending tasks in a list (for the specified queue)
	queueKey := fmt.Sprintf("asynq:{%s}:pending", queueName)

	// Fetch all task IDs from the queue (not just the first one)
	taskIDs, err := rdb.LRange(ctx, queueKey, 0, -1).Result()
	if err != nil {
		log.Printf("Error fetching tasks from queue: %v", err)
		return err
	}

	if len(taskIDs) == 0 {
		log.Println("No tasks available in the queue.")
		return nil
	}

	// Process each task
	for _, taskID := range taskIDs {
		// Fetch task details using HGETALL
		taskKey := fmt.Sprintf("asynq:{%s}:t:%s", queueName, taskID)
		taskData, err := rdb.HGetAll(ctx, taskKey).Result()
		if err != nil {
			log.Printf("Error fetching task data: %v", err)
			continue
		}

		if len(taskData) == 0 {
			log.Printf("Task data not found for task ID: %s", taskID)
			continue
		}

		// Log the task data (or perform actual processing here)
		log.Printf("Processing task: %v", taskData["msg"])

		// Remove task from the queue after processing (Dequeue)
		_, err = rdb.LRem(ctx, queueKey, 0, taskID).Result()
		if err != nil {
			log.Printf("Error removing task %s from queue: %v", taskID, err)
		} else {
			log.Printf("Task %s removed from queue", taskID)
		}
	}

	return nil
}
