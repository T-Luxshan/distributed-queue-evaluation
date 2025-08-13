package tasks

import (
	"fmt"
	"log"
	"os/exec"
)

// UserRequestPayload represents the data structure for the payload
type UserRequestPayload struct {
	UserID    string
	RequestID string
}

// EnqueueTask enqueues a job to goq
func EnqueueTask(payload UserRequestPayload) error {
	// Ensure the script has executable permissions
	//cmdChmod := exec.Command("chmod", "+x", "/home/luxshan/Documents/Luxshan/pickme/distributed-queue-evaluate/goq/job_script.sh")
	//err := cmdChmod.Run()
	//if err != nil {
	//	log.Printf("Error making script executable: %v", err)
	//	return err
	//}

	// Prepare the job command with the necessary arguments
	jobCommand := fmt.Sprintf("/home/luxshan/Documents/Luxshan/pickme/distributed-queue-evaluate/goq/job_script.sh %s %s", payload.UserID, payload.RequestID)

	// Enqueue the job using goq
	cmd := exec.Command("goq", "sub", jobCommand)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error enqueueing job: %v", err)
		return err
	}

	log.Printf("Enqueued: userID=%s, rideID=%s\n", payload.UserID, payload.RequestID)
	return nil
}
