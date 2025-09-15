package main

import (
	"errors"
	"fmt"
	"github.com/puzpuzpuz/xsync/v3"
)

func TryPeek(queue *xsync.MPMCQueueOf[UserRequestPayload]) error {
	if val, ok := queue.TryDequeue(); ok {
		fmt.Println("First value on the queue:", val)
		queue.TryEnqueue(val)
		return nil
	}
	return errors.New("queue is empty")
}
