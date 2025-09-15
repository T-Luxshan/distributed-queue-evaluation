package tasks

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// DequeueAllTasks dequeues all the tasks from the queue one by one
func DequeueAllTasks() error {
	// List the status of the jobs in the queue using the goq stat command
	cmd := exec.Command("goq", "stat")
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error fetching queue status:", err)
		return err
	}

	// Convert the output to string
	queueStatus := string(output)
	//fmt.Println("Queue Status:\n", queueStatus)

	// Look for waiting jobs and dequeue them
	lines := strings.Split(queueStatus, "\n")
	parts := strings.Fields(lines[13])
	firstJobId, err := strconv.Atoi(strings.Split(parts[3], "]")[0])
	if err != nil {
		return err
	}
	nextJobId, _ := strconv.Atoi(strings.Split(lines[9], "=")[1])

	for i := firstJobId; i < nextJobId; i++ {
		//Dequeue the job using the job ID
		cmdDequeue := exec.Command("goq", "kill", strconv.Itoa(i))
		err = cmdDequeue.Run()
		if err != nil {
			log.Printf("Error dequeuing job with ID %s: %v", i, err)
			continue
		}

		log.Printf("Dequeued task with jobID: %s\n", i)
	}

	// If no tasks were found, print a message
	fmt.Println("No tasks found to dequeue.")
	return nil
}
