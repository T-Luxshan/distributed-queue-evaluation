package main

import (
	"bufio"
	"fmt"
	"github.com/T-Luxshan/distributed-queue-evaluation/goq/tasks"
	"log"
	"os"
	"strings"
	"time"
)

//type UserRequestPayload struct {
//	UserID    string
//	RequestID string
//}

func goq() (time.Duration, time.Duration, time.Duration) {
	file, err := os.Open("../resources/ride_request_10")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1, -1, -1

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// Read line by line
	scanner := bufio.NewScanner(file)

	//var wg sync.WaitGroup
	//Measure time for enqueue
	startEnqueue := time.Now()

	for scanner.Scan() {

		line := scanner.Text()
		//wg.Add(1)
		//go func(line string) {
		//	defer wg.Done()
		request := strings.Split(line, ",") // Split by comma

		if len(request) != 3 {

			fmt.Println("Invalid record:", line)
			return -1, -1, -1
		}
		userID := request[0]
		rideID := request[1]

		payload := tasks.UserRequestPayload{UserID: userID, RequestID: rideID}

		err := tasks.EnqueueTask(payload)
		if err != nil {
			log.Println("Error enqueuing task:", err)
			return -1, -1, -1
		}

		//}(line)
	}
	//wg.Wait()
	elapsedEnqueue := time.Since(startEnqueue)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("-------------------------------------Peek the Queue-------------------------------------")

	startPeek := time.Now()
	tasks.PeekTask()

	elapsedPeek := time.Since(startPeek)

	fmt.Println("-------------------------------------Dequeue the Queue-------------------------------------")

	startDequeue := time.Now()
	if err := tasks.DequeueAllTasks(); err != nil {
		fmt.Println("Error dequeueing tasks:", err)
	}
	elapsedDequeue := time.Since(startDequeue)

	return elapsedEnqueue, elapsedPeek, elapsedDequeue
}

func main() {
	// Example server setup
	//fmt.Println(goq())
	//tasks.PeekTask()
	//if err := tasks.DequeueAllTasks(); err != nil {
	//	fmt.Println("Error dequeueing tasks:", err)

	var totalElapsedEnqueue, totalElapsedPeek, totalElapsedDequeue time.Duration
	for i := 0; i < 5; i++ {
		t1, t2, t3 := goq()
		totalElapsedEnqueue += t1
		totalElapsedPeek += t2
		totalElapsedDequeue += t3
	}
	fmt.Printf("\n\nAverage Elapsed time for Enqueue=%s\n", totalElapsedEnqueue/10)
	fmt.Printf("Average Elapsed time for Peek=%s\n", totalElapsedPeek/10)
	fmt.Printf("Average Elapsed time for Dequeue=%s\n", totalElapsedDequeue/10)
	//}
}
