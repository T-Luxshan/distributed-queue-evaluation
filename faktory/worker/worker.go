package worker

import (
	"context"
	"fmt"
	worker "github.com/contribsys/faktory_worker_go"
)

func Worker() error {
	mgr := worker.NewManager()

	mgr.Register("req", someJobWorker)
	err := mgr.Run()
	if err != nil {
		return fmt.Errorf("failed to register: %w", err)
	}
	return nil
}

func someJobWorker(ctx context.Context, args ...interface{}) error {

	if len(args) == 0 {
		return fmt.Errorf("no args")
	}
	message, ok := args[0].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid args")
	}

	userID, ok := message["userID"].(string)
	if !ok {
		return fmt.Errorf("invalid value for user id")
	}
	rideID, ok := message["requestID"].(string)
	if !ok {
		return fmt.Errorf("invalid value for ride id")
	}

	fmt.Printf("Job Executed! UserID: %s, RequestID: %s\n", userID, rideID)
	return nil
}
