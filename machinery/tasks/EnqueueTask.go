package tasks

import (
	"encoding/csv"
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	Server "github.com/T-Luxshan/distributed-queue-evaluation/machinery/server"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func Enqueue() (error, time.Duration) {
	server, err := Server.StartServer()
	if err != nil {
		return err, -1
	}

	file, err := os.Open("../resource.csv")
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err), -1
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)

	var wg sync.WaitGroup
	startEnqueue := time.Now()
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			return fmt.Errorf("failed to read CSV row: %w", err), -1
		}
		wg.Add(1)
		go func(record []string) {
			defer wg.Done()
			if len(record) != 2 {
				log.Printf("Skipping invalid row: %v (expected 3 columns)\n", record)
				return
			}

			userID := record[0]
			requestID := record[1]

			// Create a task signature.  Pass the data from the CSV row.
			signature := &tasks.Signature{
				Name: "process_tasks",
				Args: []tasks.Arg{
					{Type: "string", Value: userID},
					{Type: "string", Value: requestID},
				},
			}

			_, err = server.SendTask(signature)
			if err != nil {
				log.Printf("Failed to send task for row: %v, error: %v\n", record, err)
				return
			}

			log.Printf("Enqueued task for row: %v\n", record)
		}(record)

	}
	wg.Wait()
	elapsedEnqueue := time.Since(startEnqueue)
	fmt.Printf("Enqueued task for %s seconds\n", elapsedEnqueue)
	log.Println("All ride details enqueued.")

	return nil, elapsedEnqueue
}
