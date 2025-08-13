package tasks

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func PeekTask() {
	// List all jobs in the queue using the goq stat command
	cmd := exec.Command("goq", "stat")
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error peeking queue:", err)
		return
	}

	// Convert the output to string
	queueStatus := string(output)
	//fmt.Println("Queue Status:\n", queueStatus)

	// Look for the first job in the "waitingJobs" section
	lines := strings.Split(queueStatus, "\n")
	for _, line := range lines {
		// Check if the line contains a waiting job and the script path
		if strings.Contains(line, "wait") && strings.Contains(line, "job_script.sh") {

			parts := strings.Fields(line)
			for i, part := range parts {
				if strings.Contains(part, "job_script.sh") && i+2 < len(parts) {
					userID := parts[i+1] // userID (e.g., '001')
					rideID := parts[i+2] // rideID (e.g., 'RIDE001')

					// Print the first task's userID and rideID
					fmt.Printf("First Task : userID: %s, rideID: %s\n", userID, rideID)
					return
				}
			}
		}
	}

	// If no tasks are found in the queue, print a message
	fmt.Println("No tasks found in the queue.")
}
