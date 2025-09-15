package main

import (
	"bufio"
	"fmt"
	"github.com/puzpuzpuz/xsync/v3"
	"os"
	"strconv"
	"strings"
	"sync"
)

type UserRequestPayload struct {
	UserID    int
	RequestID string
}

func xsyncMPMCQueue() {
	file, err := os.Open("../resource.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// Create an MPMCQueue for enqueuing the requests
	queue := xsync.NewMPMCQueueOf[UserRequestPayload](60024000)

	// Read line by line
	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	for scanner.Scan() {

		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			request := strings.Split(line, ",") // Split by comma

			if len(request) != 2 {
				fmt.Println("Invalid record:", line)
				return
			}
			userID, _ := strconv.Atoi(request[0])
			requestID := request[1]

			payload := UserRequestPayload{
				UserID:    userID,
				RequestID: requestID,
			}
			//	Enqueue to the MPMCQueue
			insert := queue.TryEnqueue(payload)
			if insert {
				fmt.Printf("Inserted: %d, %s\n", payload.UserID, payload.RequestID)
			} else {
				fmt.Println("Queue is full, request dropped: ", line)
			}

		}(line)
	}
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("-------------------------------------Peek the Queue-------------------------------------")
	if err := TryPeek(queue); err != nil {
		fmt.Println("Error peeking the queue:", err)
	}

	fmt.Println("-------------------------------------Dequeue all the tasks-------------------------------------")

	wg.Add(1)
	go func() {
		defer wg.Done()
		DequeueAll(queue)
	}()
	wg.Wait()

}

func main() {
	xsyncMPMCQueue()
}
