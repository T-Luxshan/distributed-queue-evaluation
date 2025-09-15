package tasks

import (
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	"log"
	"time"
)

func dequeue(client *faktory.Client, queueName string, queueSize uint64) error {

	fmt.Printf("Dequeue from the %s Queue\n", queueName)
	if queueSize < 1 {
		return fmt.Errorf("queue is empty")
	}

	for i := 0; uint64(i) < queueSize; i++ {
		job, err := client.Fetch(queueName)
		if err != nil {
			log.Fatalf("Error fetching job: %v", err)
		}

		// Ack the job to mark as processed and dequeue from the queue.
		err = client.Ack(job.Jid)
		if err != nil {
			log.Fatalf("Error acking job: %v", err)
		}

		if len(job.Args) > 0 {
			argMap, ok := job.Args[0].(map[string]interface{})
			if !ok {
				_ = fmt.Errorf("invalid args")
			}
			userID, ok := argMap["userID"].(string)
			if !ok {
				_ = fmt.Errorf("invalid value for user id")
			}
			requestID, ok := argMap["requestID"].(string)
			if !ok {
				_ = fmt.Errorf("invalid value for request id")
			}

			fmt.Printf("Next job in the queue: UserID: %s, Request ID: %s\n", userID, requestID)
		} else {
			fmt.Println("Invalid job args")
		}
	}

	return nil
}

func DequeueTask() time.Duration {

	client, err := faktory.Open()
	if err != nil {
		log.Fatalf("Error connecting to Faktory: %v", err)
	}

	defer func(client *faktory.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(client)

	queueSizes, err := client.QueueSizes()
	if err != nil {
		log.Fatalf("Error getting queue size: %v", err)
	}

	startPeek := time.Now()
	queueTypes := []string{"critical", "default", "bulk"}
	for _, queueType := range queueTypes {
		err = dequeue(client, queueType, queueSizes[queueType])
		if err != nil {
			log.Printf("Error dequeuing from %s queue: %v", queueType, err)
		}
	}
	elapsedPeek := time.Since(startPeek)
	return elapsedPeek
}
