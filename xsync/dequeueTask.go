package main

import (
	"fmt"
	"github.com/puzpuzpuz/xsync/v3"
)

func DequeueAll(queue *xsync.MPMCQueueOf[UserRequestPayload]) {
	for {
		item, ok := queue.TryDequeue()
		if !ok {
			break // Stop when the queue is empty
		}
		fmt.Printf("Dequeued: UserID=%d, RideID=%s\n", item.UserID, item.RequestID)
	}
}
