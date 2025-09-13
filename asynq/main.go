package main

import (
	"bufio"
	"fmt"
	"github.com/T-Luxshan/distributed-queue-evaluation/asynq/tasks"
	"github.com/hibiken/asynq"
	"os"
	"strconv"
	"strings"
	"sync"
)

const redisAddr = "127.0.0.1:6379"

func asynqImp(client *asynq.Client) {

	//userID := 123
	//requestID := "req-123"

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

	// Read line by line
	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	//Measure time for enqueue
	for scanner.Scan() {

		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			request := strings.Split(line, ",")

			if len(request) != 2 {
				fmt.Println("Invalid record:", line)
				return
			}
			userID, _ := strconv.Atoi(request[0])
			requestID := request[1]
			if err != nil {
				fmt.Println("Invalid fare value:", request[2])
				return
			}
			tasks.EnqueueTask(client, userID, requestID)
		}(line)
	}
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//tasks.EnqueueTask(client, userID, requestID)
	err = tasks.PeekNextTask(redisAddr, "low")
	if err != nil {
		fmt.Println(err)
	}

	err = tasks.DequeueAll(redisAddr, "default")
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
