package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

const (
	TypeUserRequest = "user:request"
)

// UserRequestPayload Structure of the user request payload.
type UserRequestPayload struct {
	UserID    int
	RequestID string
}

// NewUserRequestTask creates a new task for a user request
func NewUserRequestTask(userID int, requestID string) (*asynq.Task, error) {
	payload, err := json.Marshal(UserRequestPayload{UserID: userID, RequestID: requestID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeUserRequest, payload), nil
}

// HandleUserRequestTask processes the user request task
func HandleUserRequestTask(ctx context.Context, t *asynq.Task) error {
	var p UserRequestPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Processing request for user %d with request ID %s", p.UserID, p.RequestID)
	return nil
}
