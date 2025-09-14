package main

import (
	"fmt"
	MyTasks "github.com/T-Luxshan/distributed-queue-evaluation/faktory/tasks"
	"time"
)

func main() {
	var elapsedEnqueue, elapsedDequeue time.Duration

	elapsedEnqueue += MyTasks.Enqueue()
	elapsedDequeue += MyTasks.DequeueTask()

	fmt.Println("\n\nAvg Enqueue:", elapsedEnqueue)
	fmt.Println("Avg Dequeue:", elapsedDequeue)

}
