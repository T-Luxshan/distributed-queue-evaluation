package tasks

import (
	"bufio"
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type UserRequestPayload struct {
	UserID    string `json:"userID"`
	RequestID string `json:"requestID"`
}

func chooseQueue() string {
	queueTypes := []string{"critical", "default", "low"}
	randomQ := queueTypes[rand.Intn(len(queueTypes))]

	return randomQ
}

func Enqueue() time.Duration {
	client, err := faktory.Open()
	if err != nil {
		panic(err)
	}
	defer func(client *faktory.Client) {
		err = client.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(client)

	file, err := os.Open("../resource.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	startEnqueue := time.Now()
	for scanner.Scan() {

		line := scanner.Text()
		record := strings.Split(line, ",")

		if len(record) != 2 {
			fmt.Println("Invalid record:", line)
			continue
		}

		request := UserRequestPayload{UserID: record[0], RequestID: record[1]}
		req := request
		job := faktory.NewJob("req", req)

		queueType := chooseQueue()
		if err != nil {
			fmt.Println("Error enqueuing job:", err)
			return -1
		}
		job.Queue = queueType
		err = client.Push(job)
		if err != nil {
			_ = fmt.Errorf("failed to push job: %w", err)
			return -1
		}
		fmt.Printf("Job enqueued to %s Queue, UserID: %s, rideID: %s\n", queueType, record[0], record[1])
	}
	elapsedEnqueue := time.Since(startEnqueue)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return elapsedEnqueue
}
